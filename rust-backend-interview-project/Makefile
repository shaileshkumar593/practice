run:
	cargo run

test:
	cargo test

check:
	cargo check

lint:
	cargo fmt --all -- --check
	cargo clippy --all-targets --all-features -- -D warnings

docker:
	docker build -t rust-backend-interview:local .

up:
	docker compose up -d

down:
	docker compose down
