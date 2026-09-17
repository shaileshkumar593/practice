# Groq AI Models - Python
Examples 01-09 cover LLM, reasoning, embeddings, reranking, multimodal input,
tool calling, agents, RAG, and RAG+rering.

Setup:
python -m venv .venv
pip install -r requirements.txt
copy .env.example .env

Groq is used for generation/reasoning. Embeddings and reranking are separate
components. In production, authorization, retrieval and tool execution remain
application responsibilities.
