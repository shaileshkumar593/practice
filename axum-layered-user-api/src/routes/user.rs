use crate::{
    handlers::user::{create_user, delete_user, get_user, update_user},
    state::AppState,
};
use axum::{
    routing::{delete, get, post, put},
    Router,
};
use std::sync::Arc;

pub fn user_routes() -> Router<Arc<AppState>> {
    Router::new()
        .route("/users", post(create_user))
        .route("/users/{id}", get(get_user))
        .route("/users/{id}", put(update_user))
        .route("/users/{id}", delete(delete_user))
}
