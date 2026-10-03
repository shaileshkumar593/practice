.PHONY: install format lint test migrate run

install:
	pip install -r requirements.txt

format:
	ruff format app tests

lint:
	ruff check app tests

test:
	pytest -q

migrate:
	alembic upgrade head

run:
	uvicorn app.main:app --host 0.0.0.0 --port 8000
