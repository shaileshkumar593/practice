use axum::{
    extract::{Path, State},
    middleware,
    routing::{get, post},
    Json, Router,
};
use uuid::Uuid;
use validator::Validate;

use crate::{
    error::AppError,
    middleware::request_id,
    models::{CreateUserRequest, TransferRequest, User},
    AppState,
};

pub fn router() -> Router<AppState> {
    Router::new()
        .route("/users", post(create_user))
        .route("/users/{id}", get(get_user))
        .route("/transfers", post(transfer))
        .layer(middleware::from_fn(request_id))
}

async fn create_user(
    State(state): State<AppState>,
    Json(req): Json<CreateUserRequest>,
) -> Result<Json<User>, AppError> {
    req.validate()
        .map_err(|e| AppError::Validation(e.to_string()))?;

    let user = state.services.create_user(req).await?;
    Ok(Json(user))
}

async fn get_user(
    State(state): State<AppState>,
    Path(id): Path<Uuid>,
) -> Result<Json<User>, AppError> {
    Ok(Json(state.services.get_user(id).await?))
}

async fn transfer(
    State(state): State<AppState>,
    Json(req): Json<TransferRequest>,
) -> Result<Json<serde_json::Value>, AppError> {
    if req.amount <= 0 {
        return Err(AppError::Validation("amount must be positive".into()));
    }

    state.services.transfer(req).await?;
    Ok(Json(serde_json::json!({"status": "accepted"})))
}
