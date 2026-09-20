use crate::{
    errors::AppError,
    models::user::{CreateUserRequest, UpdateUserRequest, User},
    repositories::user_repository::UserRepository,
};
use uuid::Uuid;

#[derive(Clone)]
pub struct UserService<R> {
    repository: R,
}

impl UserService<UserRepository> {
    pub fn new(repository: UserRepository) -> Self {
        Self { repository }
    }

    pub fn create_user(&self, request: CreateUserRequest) -> Result<User, AppError> {
        validate(&request.name, &request.email)?;

        let user = User {
            id: Uuid::new_v4(),
            name: request.name,
            email: request.email,
        };

        self.repository.create(user)
    }

    pub fn get_user(&self, id: Uuid) -> Result<User, AppError> {
        self.repository.find_by_id(id)
    }

    pub fn update_user(
        &self,
        id: Uuid,
        request: UpdateUserRequest,
    ) -> Result<User, AppError> {
        validate(&request.name, &request.email)?;

        let user = User {
            id,
            name: request.name,
            email: request.email,
        };

        self.repository.update(user)
    }

    pub fn delete_user(&self, id: Uuid) -> Result<(), AppError> {
        self.repository.delete(id)
    }
}

fn validate(name: &str, email: &str) -> Result<(), AppError> {
    if name.trim().is_empty() {
        return Err(AppError::Validation("name cannot be empty".into()));
    }

    if !email.contains('@') {
        return Err(AppError::Validation("invalid email".into()));
    }

    Ok(())
}
