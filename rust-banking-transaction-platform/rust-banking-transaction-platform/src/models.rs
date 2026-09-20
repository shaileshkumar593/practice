use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Serialize)]
pub struct User {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize)]
pub struct CreateUserRequest {
    pub name: String,
    pub email: String,
}

#[derive(Debug, Serialize)]
pub struct Account {
    pub id: Uuid,
    pub user_id: Uuid,
    pub balance: i64,
    pub currency: String,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize)]
pub struct CreateAccountRequest {
    pub user_id: Uuid,
    pub currency: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct TransferRequest {
    pub from_account_id: Uuid,
    pub to_account_id: Uuid,
    pub amount: i64,
    pub currency: String,
    pub idempotency_key: String,
}

#[derive(Debug, Serialize)]
pub struct TransferResponse {
    pub transaction_id: Uuid,
    pub status: String,
}

#[derive(Debug, Serialize)]
pub struct HealthResponse {
    pub status: &'static str,
}
