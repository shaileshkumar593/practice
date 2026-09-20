use crate::{
    error::AppError,
    models::{
        Account, CreateAccountRequest, CreateUserRequest, HealthResponse, TransferRequest,
        TransferResponse, User,
    },
    state::AppState,
};
use axum::{
    extract::{Path, State},
    routing::{get, post},
    Json, Router,
};
use chrono::Utc;
use sqlx::Row;
use uuid::Uuid;

pub fn router() -> Router<AppState> {
    Router::new()
        .route("/health", get(health))
        .route("/users", post(create_user))
        .route("/users/{id}", get(get_user))
        .route("/accounts", post(create_account))
        .route("/accounts/{id}", get(get_account))
        .route("/transfers", post(create_transfer))
}

async fn health() -> Json<HealthResponse> {
    Json(HealthResponse { status: "ok" })
}

async fn create_user(
    State(state): State<AppState>,
    Json(req): Json<CreateUserRequest>,
) -> Result<Json<User>, AppError> {
    if req.name.trim().is_empty() || req.email.trim().is_empty() {
        return Err(AppError::Validation("name and email are required".into()));
    }

    let id = Uuid::new_v4();
    let created_at = Utc::now();

    sqlx::query(
        "INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)",
    )
    .bind(id)
    .bind(&req.name)
    .bind(&req.email)
    .bind(created_at)
    .execute(&state.db)
    .await?;

    Ok(Json(User {
        id,
        name: req.name,
        email: req.email,
        created_at,
    }))
}

async fn get_user(
    State(state): State<AppState>,
    Path(id): Path<Uuid>,
) -> Result<Json<User>, AppError> {
    let row = sqlx::query(
        "SELECT id, name, email, created_at FROM users WHERE id = $1",
    )
    .bind(id)
    .fetch_optional(&state.db)
    .await?
    .ok_or(AppError::NotFound)?;

    Ok(Json(User {
        id: row.try_get("id")?,
        name: row.try_get("name")?,
        email: row.try_get("email")?,
        created_at: row.try_get("created_at")?,
    }))
}

async fn create_account(
    State(state): State<AppState>,
    Json(req): Json<CreateAccountRequest>,
) -> Result<Json<Account>, AppError> {
    let currency = req.currency.unwrap_or_else(|| "INR".into());
    let id = Uuid::new_v4();
    let created_at = Utc::now();

    sqlx::query(
        "INSERT INTO accounts (id, user_id, balance, currency, created_at)
         VALUES ($1, $2, 0, $3, $4)",
    )
    .bind(id)
    .bind(req.user_id)
    .bind(&currency)
    .bind(created_at)
    .execute(&state.db)
    .await?;

    Ok(Json(Account {
        id,
        user_id: req.user_id,
        balance: 0,
        currency,
        created_at,
    }))
}

async fn get_account(
    State(state): State<AppState>,
    Path(id): Path<Uuid>,
) -> Result<Json<Account>, AppError> {
    let row = sqlx::query(
        "SELECT id, user_id, balance, currency, created_at
         FROM accounts WHERE id = $1",
    )
    .bind(id)
    .fetch_optional(&state.db)
    .await?
    .ok_or(AppError::NotFound)?;

    Ok(Json(Account {
        id: row.try_get("id")?,
        user_id: row.try_get("user_id")?,
        balance: row.try_get("balance")?,
        currency: row.try_get("currency")?,
        created_at: row.try_get("created_at")?,
    }))
}

async fn create_transfer(
    State(state): State<AppState>,
    Json(req): Json<TransferRequest>,
) -> Result<Json<TransferResponse>, AppError> {
    if req.amount <= 0 {
        return Err(AppError::Validation("amount must be positive".into()));
    }

    if req.from_account_id == req.to_account_id {
        return Err(AppError::Validation("accounts must be different".into()));
    }

    let mut tx = state.db.begin().await?;

    // Idempotency: if this key was already processed, return the original transaction.
    let existing = sqlx::query(
        "SELECT transaction_id FROM idempotency_keys WHERE idempotency_key = $1",
    )
    .bind(&req.idempotency_key)
    .fetch_optional(&mut *tx)
    .await?;

    if let Some(row) = existing {
        let transaction_id: Uuid = row.try_get("transaction_id")?;
        tx.commit().await?;
        return Ok(Json(TransferResponse {
            transaction_id,
            status: "already_processed".into(),
        }));
    }

    // Lock accounts in deterministic order to reduce deadlock risk.
    let (first, second) = if req.from_account_id < req.to_account_id {
        (req.from_account_id, req.to_account_id)
    } else {
        (req.to_account_id, req.from_account_id)
    };

    let first_row = sqlx::query(
        "SELECT id, balance, currency FROM accounts WHERE id = $1 FOR UPDATE",
    )
    .bind(first)
    .fetch_optional(&mut *tx)
    .await?
    .ok_or(AppError::NotFound)?;

    let second_row = sqlx::query(
        "SELECT id, balance, currency FROM accounts WHERE id = $1 FOR UPDATE",
    )
    .bind(second)
    .fetch_optional(&mut *tx)
    .await?
    .ok_or(AppError::NotFound)?;

    let first_currency: String = first_row.try_get("currency")?;
    let second_currency: String = second_row.try_get("currency")?;

    if first_currency != req.currency || second_currency != req.currency {
        return Err(AppError::Validation("currency mismatch".into()));
    }

    let from_row = if first == req.from_account_id {
        first_row
    } else {
        second_row
    };

    let from_balance: i64 = from_row.try_get("balance")?;

    if from_balance < req.amount {
        return Err(AppError::Conflict("insufficient funds".into()));
    }

    sqlx::query(
        "UPDATE accounts
         SET balance = balance - $1
         WHERE id = $2",
    )
    .bind(req.amount)
    .bind(req.from_account_id)
    .execute(&mut *tx)
    .await?;

    sqlx::query(
        "UPDATE accounts
         SET balance = balance + $1
         WHERE id = $2",
    )
    .bind(req.amount)
    .bind(req.to_account_id)
    .execute(&mut *tx)
    .await?;

    let transaction_id = Uuid::new_v4();

    sqlx::query(
        "INSERT INTO transactions
         (id, from_account_id, to_account_id, amount, currency, status, created_at)
         VALUES ($1, $2, $3, $4, $5, 'COMPLETED', $6)",
    )
    .bind(transaction_id)
    .bind(req.from_account_id)
    .bind(req.to_account_id)
    .bind(req.amount)
    .bind(&req.currency)
    .bind(Utc::now())
    .execute(&mut *tx)
    .await?;

    // Transactional outbox: database state + event are committed together.
    let event_payload = serde_json::json!({
        "transaction_id": transaction_id,
        "from_account_id": req.from_account_id,
        "to_account_id": req.to_account_id,
        "amount": req.amount,
        "currency": req.currency,
        "event_type": "TransferCompleted"
    });

    sqlx::query(
        "INSERT INTO outbox_events (id, aggregate_id, event_type, payload, created_at)
         VALUES ($1, $2, $3, $4, $5)",
    )
    .bind(Uuid::new_v4())
    .bind(transaction_id)
    .bind("TransferCompleted")
    .bind(event_payload)
    .bind(Utc::now())
    .execute(&mut *tx)
    .await?;

    sqlx::query(
        "INSERT INTO idempotency_keys (idempotency_key, transaction_id, created_at)
         VALUES ($1, $2, $3)",
    )
    .bind(&req.idempotency_key)
    .bind(transaction_id)
    .bind(Utc::now())
    .execute(&mut *tx)
    .await?;

    tx.commit().await?;

    Ok(Json(TransferResponse {
        transaction_id,
        status: "completed".into(),
    }))
}
