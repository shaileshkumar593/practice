use chrono::Utc;
use serde_json::json;
use sqlx::{PgPool, Postgres, Transaction};
use uuid::Uuid;

use crate::{
    error::AppError,
    kafka::KafkaProducer,
    models::{CreateUserRequest, Event, TransferRequest, User},
    redis_cache::RedisCache,
    repositories,
};

pub struct AppServices {
    pub db: PgPool,
    pub redis: RedisCache,
    pub kafka: KafkaProducer,
    pub jwt_secret: String,
}

impl AppServices {
    pub fn new(db: PgPool, redis: RedisCache, kafka: KafkaProducer, jwt_secret: String) -> Self {
        Self { db, redis, kafka, jwt_secret }
    }

    pub async fn create_user(&self, req: CreateUserRequest) -> Result<User, AppError> {
        repositories::create_user(&self.db, req).await
    }

    pub async fn get_user(&self, id: Uuid) -> Result<User, AppError> {
        let key = format!("user:{id}");

        if let Some(cached) = self.redis.get(&key).await? {
            return serde_json::from_str(&cached)
                .map_err(|e| AppError::Internal(e.to_string()));
        }

        let user = repositories::get_user(&self.db, id).await?;
        let payload = serde_json::to_string(&user)
            .map_err(|e| AppError::Internal(e.to_string()))?;

        self.redis.set_ex(&key, &payload, 300).await?;
        Ok(user)
    }

    pub async fn transfer(&self, req: TransferRequest) -> Result<(), AppError> {
        let mut tx: Transaction<'_, Postgres> = self.db.begin().await?;

        let debit = sqlx::query(
            "UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1"
        )
        .bind(req.amount)
        .bind(req.from_user)
        .execute(&mut *tx)
        .await?;

        if debit.rows_affected() != 1 {
            tx.rollback().await?;
            return Err(AppError::Conflict("insufficient funds or account missing".into()));
        }

        sqlx::query("UPDATE accounts SET balance = balance + $1 WHERE id = $2")
            .bind(req.amount)
            .bind(req.to_user)
            .execute(&mut *tx)
            .await?;

        let event = Event {
            event_type: "money.transferred".to_string(),
            event_id: Uuid::new_v4(),
            payload: json!({
                "from_user": req.from_user,
                "to_user": req.to_user,
                "amount": req.amount,
                "created_at": Utc::now(),
            }),
        };

        sqlx::query(
            "INSERT INTO outbox (id, event_type, payload, created_at) VALUES ($1, $2, $3, NOW())"
        )
        .bind(event.event_id)
        .bind(&event.event_type)
        .bind(&event.payload)
        .execute(&mut *tx)
        .await?;

        tx.commit().await?;
        Ok(())
    }
}
