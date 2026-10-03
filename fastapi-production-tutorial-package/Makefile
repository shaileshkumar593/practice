install:
	uv sync

dev:
	uv run fastapi dev app/main.py

run:
	uv run fastapi run app/main.py --host 0.0.0.0 --port 8000

test:
	uv run pytest -q

lint:
	uv run ruff check .

format:
	uv run ruff format .

docker-build:
	docker build -t fastapi-production-tutorial:local .

docker-up:
	docker compose up --build

docker-down:
	docker compose down
