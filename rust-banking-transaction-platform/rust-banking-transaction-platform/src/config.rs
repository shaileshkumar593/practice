use std::env;

#[derive(Clone, Debug)]
pub struct Config {
    pub port: u16,
    pub database_url: String,
    pub redis_url: String,
    pub request_timeout_secs: u64,
}

impl Config {
    pub fn from_env() -> Result<Self, Box<dyn std::error::Error>> {
        Ok(Self {
            port: env::var("PORT").unwrap_or_else(|_| "3000".into()).parse()?,
            database_url: env::var("DATABASE_URL")
                .unwrap_or_else(|_| "postgres://banking:banking@localhost:5432/banking".into()),
            redis_url: env::var("REDIS_URL")
                .unwrap_or_else(|_| "redis://127.0.0.1:6379".into()),
            request_timeout_secs: env::var("REQUEST_TIMEOUT_SECS")
                .unwrap_or_else(|_| "10".into())
                .parse()?,
        })
    }
}
