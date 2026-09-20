use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;
use validator::Validate;

#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct User {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize, Validate)]
pub struct CreateUserRequest {
    #[validate(length(min = 2, max = 100))]
    pub name: String,
    #[validate(email)]
    pub email: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct TransferRequest {
    pub from_user: Uuid,
    pub to_user: Uuid,
    pub amount: i64,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Event<T> {
    pub event_type: String,
    pub event_id: Uuid,
    pub payload: T,
}
