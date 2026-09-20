use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
    Json,
};
use serde::Serialize;
use thiserror::Error;

#[derive(Debug, Error)]
pub enum AppError {
    #[error("validation error: {0}")]
    Validation(String),

    #[error("resource not found")]
    NotFound,

    #[error("conflict: {0}")]
    Conflict(String),

    #[error("database error: {0}")]
    Database(#[from] sqlx::Error),

    #[error("redis error: {0}")]
    Redis(#[from] redis::RedisError),

    #[error("internal server error")]
    Internal,
}

#[derive(Serialize)]
struct ErrorBody {
    error: String,
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, message) = match self {
            Self::Validation(msg) => (StatusCode::BAD_REQUEST, msg),
            Self::NotFound => (StatusCode::NOT_FOUND, "resource not found".into()),
            Self::Conflict(msg) => (StatusCode::CONFLICT, msg),
            Self::Database(_) | Self::Redis(_) | Self::Internal => {
                (StatusCode::INTERNAL_SERVER_ERROR, "internal server error".into())
            }
        };

        (status, Json(ErrorBody { error: message })).into_response()
    }
}
