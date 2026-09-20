use crate::{
    errors::AppError,
    models::user::{CreateUserRequest, UpdateUserRequest, User},
    state::AppState,
};
use axum::{
    extract::{Path, State},
    http::StatusCode,
    Json,
};
use std::sync::Arc;
use uuid::Uuid;

pub async fn create_user(
    State(state): State<Arc<AppState>>,
    Json(request): Json<CreateUserRequest>,
) -> Result<(StatusCode, Json<User>), AppError> {
    let user = state.user_service.create_user(request)?;
    Ok((StatusCode::CREATED, Json(user)))
}

pub async fn get_user(
    State(state): State<Arc<AppState>>,
    Path(id): Path<Uuid>,
) -> Result<Json<User>, AppError> {
    Ok(Json(state.user_service.get_user(id)?))
}

pub async fn update_user(
    State(state): State<Arc<AppState>>,
    Path(id): Path<Uuid>,
    Json(request): Json<UpdateUserRequest>,
) -> Result<Json<User>, AppError> {
    Ok(Json(state.user_service.update_user(id, request)?))
}

pub async fn delete_user(
    State(state): State<Arc<AppState>>,
    Path(id): Path<Uuid>,
) -> Result<StatusCode, AppError> {
    state.user_service.delete_user(id)?;
    Ok(StatusCode::NO_CONTENT)
}
