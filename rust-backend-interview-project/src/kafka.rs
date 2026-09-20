use crate::error::AppError;
use rdkafka::{
    producer::{FutureProducer, FutureRecord},
    ClientConfig,
};
use std::time::Duration;

#[derive(Clone)]
pub struct KafkaProducer {
    producer: FutureProducer,
    topic: String,
}

impl KafkaProducer {
    pub fn new(brokers: &str, topic: &str) -> Result<Self, AppError> {
        let producer = ClientConfig::new()
            .set("bootstrap.servers", brokers)
            .set("message.timeout.ms", "5000")
            .create()
            .map_err(|e| AppError::Kafka(e.to_string()))?;

        Ok(Self {
            producer,
            topic: topic.to_string(),
        })
    }

    pub async fn publish(&self, key: &str, payload: String) -> Result<(), AppError> {
        self.producer
            .send(
                FutureRecord::to(&self.topic)
                    .key(key)
                    .payload(&payload),
                Duration::from_secs(5),
            )
            .await
            .map_err(|(e, _)| AppError::Kafka(e.to_string()))?;
        Ok(())
    }
}
