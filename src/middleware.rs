use axum::{extract::Request, middleware::Next, response::Response};
use tracing::info;
use uuid::Uuid;

pub async fn request_id(mut request: Request, next: Next) -> Response {
    let id = Uuid::new_v4().to_string();
    request.headers_mut().insert(
        "x-request-id",
        id.parse().expect("valid header"),
    );
    info!(request_id = %id, method = %request.method(), uri = %request.uri(), "request");
    next.run(request).await
}
