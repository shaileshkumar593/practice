use crate::error::AppError;
use redis::{aio::ConnectionManager, AsyncCommands};

#[derive(Clone)]
pub struct RedisCache {
    manager: ConnectionManager,
}

impl RedisCache {
    pub async fn new(url: &str) -> Result<Self, AppError> {
        let client = redis::Client::open(url)
            .map_err(|e| AppError::Redis(e.to_string()))?;
        let manager = client
            .get_connection_manager()
            .await
            .map_err(|e| AppError::Redis(e.to_string()))?;
        Ok(Self { manager })
    }

    pub async fn get(&self, key: &str) -> Result<Option<String>, AppError> {
        let mut conn = self.manager.clone();
        conn.get(key).await.map_err(|e| AppError::Redis(e.to_string()))
    }

    pub async fn set_ex(&self, key: &str, value: &str, ttl: u64) -> Result<(), AppError> {
        let mut conn = self.manager.clone();
        conn.set_ex::<_, _, ()>(key, value, ttl)
            .await
            .map_err(|e| AppError::Redis(e.to_string()))
    }
}
