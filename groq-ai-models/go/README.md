# Groq AI Models - Go
The Go project uses Groq's OpenAI-compatible Chat Completions HTTP API.

Run:
copy .env.example .env
go mod tidy
go run .

The code demonstrates LLM, reasoning and tool calling. Embeddings, reranking
and multimodal processing are shown as production architecture boundaries:
use a dedicated embedding/reranking/multimodal provider or local models as
appropriate.

For production add timeouts, controlled retries, tool authorization, input
validation, agent step limits, tracing, metrics, secret management and
tenant-aware retrieval authorization.
