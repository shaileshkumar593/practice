mod config;
mod errors;
mod handlers;
mod models;
mod repositories;
mod routes;
mod services;
mod state;

use axum::Router;
use config::Config;
use routes::user::user_routes;
use state::AppState;
use std::sync::Arc;

#[tokio::main]
async fn main() {
    let config = Config::from_env();
    let state = Arc::new(AppState::new());

    let app = Router::new()
        .merge(user_routes())
        .with_state(state);

    let listener = tokio::net::TcpListener::bind(&config.server_addr)
        .await
        .expect("failed to bind server");

    println!("Server running at http://{}", config.server_addr);

    axum::serve(listener, app).await.expect("server error");
}
