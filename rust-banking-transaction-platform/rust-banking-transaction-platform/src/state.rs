use crate::config::Config;
use sqlx::{postgres::PgPoolOptions, PgPool};
use redis::Client;

#[derive(Clone)]
pub struct AppState {
    pub db: PgPool,
    pub redis: Client,
}

impl AppState {
    pub async fn new(config: &Config) -> Result<Self, Box<dyn std::error::Error>> {
        let db = PgPoolOptions::new()
            .max_connections(20)
            .connect(&config.database_url)
            .await?;

        let redis = Client::open(config.redis_url.clone())?;

        Ok(Self { db, redis })
    }
}
