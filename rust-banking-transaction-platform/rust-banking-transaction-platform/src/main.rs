mod config;
mod error;
mod models;
mod routes;
mod state;

use axum::Router;
use config::Config;
use state::AppState;
use std::net::SocketAddr;
use tokio::net::TcpListener;
use tower_http::{cors::CorsLayer, timeout::TimeoutLayer, trace::TraceLayer};
use tracing::info;
use std::time::Duration;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenvy::dotenv().ok();

    tracing_subscriber::fmt()
        .with_env_filter(
            std::env::var("RUST_LOG")
                .unwrap_or_else(|_| "rust_banking_transaction_platform=info,tower_http=info".into()),
        )
        .json()
        .init();

    let config = Config::from_env()?;
    let state = AppState::new(&config).await?;

    let app = Router::new()
        .merge(routes::router())
        .with_state(state)
        .layer(TraceLayer::new_for_http())
        .layer(TimeoutLayer::new(Duration::from_secs(config.request_timeout_secs)))
        .layer(CorsLayer::permissive());

    let addr: SocketAddr = format!("0.0.0.0:{}", config.port).parse()?;
    let listener = TcpListener::bind(addr).await?;
    info!(%addr, "server started");

    axum::serve(listener, app).await?;
    Ok(())
}
