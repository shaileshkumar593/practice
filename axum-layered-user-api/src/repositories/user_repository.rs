use crate::{errors::AppError, models::user::User};
use std::collections::HashMap;
use std::sync::{Arc, RwLock};
use uuid::Uuid;

#[derive(Clone)]
pub struct UserRepository {
    users: Arc<RwLock<HashMap<Uuid, User>>>,
}

impl UserRepository {
    pub fn new() -> Self {
        Self {
            users: Arc::new(RwLock::new(HashMap::new())),
        }
    }

    pub fn create(&self, user: User) -> Result<User, AppError> {
        let mut users = self.users.write().map_err(|_| AppError::Internal)?;
        users.insert(user.id, user.clone());
        Ok(user)
    }

    pub fn find_by_id(&self, id: Uuid) -> Result<User, AppError> {
        let users = self.users.read().map_err(|_| AppError::Internal)?;
        users.get(&id).cloned().ok_or(AppError::NotFound)
    }

    pub fn update(&self, user: User) -> Result<User, AppError> {
        let mut users = self.users.write().map_err(|_| AppError::Internal)?;

        if !users.contains_key(&user.id) {
            return Err(AppError::NotFound);
        }

        users.insert(user.id, user.clone());
        Ok(user)
    }

    pub fn delete(&self, id: Uuid) -> Result<(), AppError> {
        let mut users = self.users.write().map_err(|_| AppError::Internal)?;

        if users.remove(&id).is_none() {
            return Err(AppError::NotFound);
        }

        Ok(())
    }
}
