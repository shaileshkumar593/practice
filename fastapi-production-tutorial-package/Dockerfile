FROM python:3.12-slim
ENV PYTHONDONTWRITEBYTECODE=1 PYTHONUNBUFFERED=1 PYTHONPATH=/app
WORKDIR /app
RUN useradd --create-home --uid 10001 appuser
RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir "fastapi[standard]" "uvicorn[standard]" "sqlalchemy[asyncio]" asyncpg pydantic-settings "pyjwt[crypto]" "pwdlib[argon2]" python-multipart sse-starlette jinja2
COPY app ./app
COPY pyproject.toml ./
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8000
CMD ["fastapi","run","app/main.py","--host","0.0.0.0","--port","8000"]
