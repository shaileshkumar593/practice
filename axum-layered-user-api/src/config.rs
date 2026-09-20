#[derive(Clone, Debug)]
pub struct Config {
    pub server_addr: String,
}

impl Config {
    pub fn from_env() -> Self {
        let port = std::env::var("PORT").unwrap_or_else(|_| "3000".to_string());

        Self {
            server_addr: format!("0.0.0.0:{port}"),
        }
    }
}
