use crate::repositories::user_repository::UserRepository;
use crate::services::user_service::UserService;
use std::sync::Arc;

#[derive(Clone)]
pub struct AppState {
    pub user_service: Arc<UserService<UserRepository>>,
}

impl AppState {
    pub fn new() -> Self {
        let repository = UserRepository::new();
        let service = UserService::new(repository);

        Self {
            user_service: Arc::new(service),
        }
    }
}
