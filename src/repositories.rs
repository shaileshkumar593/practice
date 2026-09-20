use sqlx::PgPool;
use uuid::Uuid;

use crate::{error::AppError, models::{CreateUserRequest, User}};

pub async fn init_db(pool: &PgPool) -> Result<(), sqlx::Error> {
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS users (
            id UUID PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            created_at TIMESTAMPTZ NOT NULL
        );

        CREATE TABLE IF NOT EXISTS accounts (
            id UUID PRIMARY KEY,
            user_id UUID NOT NULL REFERENCES users(id),
            balance BIGINT NOT NULL DEFAULT 0,
            created_at TIMESTAMPTZ NOT NULL
        );

        CREATE TABLE IF NOT EXISTS outbox (
            id UUID PRIMARY KEY,
            event_type TEXT NOT NULL,
            payload JSONB NOT NULL,
            created_at TIMESTAMPTZ NOT NULL,
            published_at TIMESTAMPTZ
        );
        "#,
    )
    .execute(pool)
    .await?;
    Ok(())
}

pub async fn create_user(pool: &PgPool, req: CreateUserRequest) -> Result<User, AppError> {
    let user = sqlx::query_as::<_, User>(
        r#"
        INSERT INTO users (id, name, email, created_at)
        VALUES ($1, $2, $3, NOW())
        RETURNING id, name, email, created_at
        "#,
    )
    .bind(Uuid::new_v4())
    .bind(req.name)
    .bind(req.email)
    .fetch_one(pool)
    .await?;

    Ok(user)
}

pub async fn get_user(pool: &PgPool, id: Uuid) -> Result<User, AppError> {
    sqlx::query_as::<_, User>(
        "SELECT id, name, email, created_at FROM users WHERE id = $1"
    )
    .bind(id)
    .fetch_optional(pool)
    .await?
    .ok_or(AppError::NotFound)
}
