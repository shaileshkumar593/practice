# Complete AI Model Examples in Python

Includes complete examples for:

1. LLM
2. Reasoning model
3. Embedding model
4. Reranker
5. Multimodal model
6. Tool calling
7. Agent
8. RAG
9. RAG + reranking

Setup:

```bash
python -m venv .venv
# Windows
.venv\Scripts\activate
# Linux/macOS
source .venv/bin/activate

pip install -r requirements.txt
copy .env.example .env
```

Set OPENAI_API_KEY.

Run:

```bash
python 01_llm.py
python 02_reasoning_model.py
python 03_embedding_model.py
python 04_reranker.py
python 05_multimodal.py https://example.com/image.jpg
python 06_tool_calling.py
python 07_agent.py
python 08_rag.py
python 09_combined_rag_reranker.py
```

Architecture:

```text
LLM            -> Generate
Reasoning      -> Complex multi-step solving
Embedding      -> Text to vectors
Reranker       -> Candidate relevance ranking
Multimodal     -> Text/image/etc understanding
Tool calling   -> Structured external function calls
Agent          -> Model + tools + state + bounded loop
RAG            -> Retrieve + rerank + generate
```

Production security:
- Never expose API keys.
- Validate every tool argument.
- Authorize tools independently.
- Treat retrieved content as untrusted.
- Defend against prompt injection.
- Bound agent steps, cost and tool permissions.
- Do not expose private chain-of-thought.
