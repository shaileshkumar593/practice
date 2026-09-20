use anyhow::{Context, Result};

#[derive(Clone)]
pub struct Config {
    pub database_url: String,
    pub redis_url: String,
    pub kafka_brokers: String,
    pub kafka_topic: String,
    pub jwt_secret: String,
    pub port: u16,
}

impl Config {
    pub fn from_env() -> Result<Self> {
        Ok(Self {
            database_url: std::env::var("DATABASE_URL").context("DATABASE_URL missing")?,
            redis_url: std::env::var("REDIS_URL").unwrap_or_else(|_| "redis://localhost:6379".into()),
            kafka_brokers: std::env::var("KAFKA_BROKERS").unwrap_or_else(|_| "localhost:9092".into()),
            kafka_topic: std::env::var("KAFKA_TOPIC").unwrap_or_else(|_| "banking.events".into()),
            jwt_secret: std::env::var("JWT_SECRET").context("JWT_SECRET missing")?,
            port: std::env::var("APP_PORT").unwrap_or_else(|_| "3000".into()).parse()?,
        })
    }
}
