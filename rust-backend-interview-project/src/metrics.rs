use anyhow::Result;
use once_cell::sync::Lazy;
use prometheus::{Encoder, HistogramVec, IntCounterVec, Opts, Registry, TextEncoder};

static REGISTRY: Lazy<Registry> = Lazy::new(Registry::new);

pub struct Metrics {
    pub requests: IntCounterVec,
    pub latency: HistogramVec,
}

impl Metrics {
    pub fn new() -> Result<Self> {
        let requests = IntCounterVec::new(
            Opts::new("http_requests_total", "Total HTTP requests"),
            &["method", "path", "status"],
        )?;
        let latency = HistogramVec::new(
            prometheus::HistogramOpts::new("http_request_duration_seconds", "HTTP latency"),
            &["method", "path"],
        )?;
        REGISTRY.register(Box::new(requests.clone()))?;
        REGISTRY.register(Box::new(latency.clone()))?;
        Ok(Self { requests, latency })
    }

    pub fn render(&self) -> String {
        let mut buffer = Vec::new();
        TextEncoder::new().encode(&REGISTRY.gather(), &mut buffer).unwrap();
        String::from_utf8(buffer).unwrap_or_default()
    }
}
