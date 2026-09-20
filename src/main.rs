use axum::{extract::State, http::StatusCode, routing::get, Json, Router};
use serde::Serialize;
use sqlx::PgPool;
use std::{net::SocketAddr, sync::Arc};
use tokio::net::TcpListener;
use tower_http::{cors::CorsLayer, trace::TraceLayer};
use tracing::info;

mod config;
mod error;
mod models;
mod metrics;
mod middleware;
mod handlers;
mod repositories;
mod services;
mod kafka;
mod redis_cache;

use config::Config;
use metrics::Metrics;
use services::AppServices;

#[derive(Clone)]
pub struct AppState {
    pub services: Arc<AppServices>,
    pub metrics: Arc<Metrics>,
}

#[derive(Serialize)]
struct Health {
    status: &'static str,
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenvy::dotenv().ok();

    tracing_subscriber::fmt()
        .with_env_filter(
            std::env::var("RUST_LOG")
                .unwrap_or_else(|_| "info,rust_backend_interview=debug".into()),
        )
        .json()
        .init();

    let config = Config::from_env()?;
    let pool = PgPool::connect(&config.database_url).await?;

    repositories::init_db(&pool).await?;

    let redis = redis_cache::RedisCache::new(&config.redis_url).await?;
    let kafka = kafka::KafkaProducer::new(&config.kafka_brokers, &config.kafka_topic)?;

    let services = AppServices::new(pool, redis, kafka, config.jwt_secret.clone());
    let metrics = Arc::new(Metrics::new()?);

    let state = AppState {
        services: Arc::new(services),
        metrics,
    };

    let app = Router::new()
        .route("/health", get(health))
        .route("/ready", get(ready))
        .route("/metrics", get(metrics))
        .nest("/api/v1", handlers::router())
        .layer(CorsLayer::permissive())
        .layer(TraceLayer::new_for_http())
        .with_state(state.clone());

    let addr = SocketAddr::from(([0, 0, 0, 0], config.port));
    let listener = TcpListener::bind(addr).await?;
    info!(%addr, "server started");

    axum::serve(listener, app).await?;
    Ok(())
}

async fn health() -> Json<Health> {
    Json(Health { status: "ok" })
}

async fn ready(State(state): State<AppState>) -> Result<Json<Health>, StatusCode> {
    sqlx::query("SELECT 1")
        .execute(&state.services.db)
        .await
        .map_err(|_| StatusCode::SERVICE_UNAVAILABLE)?;

    Ok(Json(Health { status: "ready" }))
}

async fn metrics(State(state): State<AppState>) -> Result<String, StatusCode> {
    Ok(state.metrics.render())
}
