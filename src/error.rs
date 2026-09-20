use axum::{http::StatusCode, response::{IntoResponse, Response}, Json};
use serde::Serialize;
use thiserror::Error;

#[derive(Debug, Error)]
pub enum AppError {
    #[error("not found")]
    NotFound,
    #[error("conflict: {0}")]
    Conflict(String),
    #[error("validation error: {0}")]
    Validation(String),
    #[error("unauthorized")]
    Unauthorized,
    #[error("database error")]
    Database(#[from] sqlx::Error),
    #[error("redis error: {0}")]
    Redis(String),
    #[error("kafka error: {0}")]
    Kafka(String),
    #[error("internal error: {0}")]
    Internal(String),
}

#[derive(Serialize)]
struct ErrorBody {
    error: String,
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let status = match self {
            Self::NotFound => StatusCode::NOT_FOUND,
            Self::Conflict(_) => StatusCode::CONFLICT,
            Self::Validation(_) => StatusCode::BAD_REQUEST,
            Self::Unauthorized => StatusCode::UNAUTHORIZED,
            Self::Database(_) | Self::Redis(_) | Self::Kafka(_) | Self::Internal(_) =>
                StatusCode::INTERNAL_SERVER_ERROR,
        };

        (status, Json(ErrorBody { error: self.to_string() })).into_response()
    }
}
