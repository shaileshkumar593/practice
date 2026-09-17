import os
from dotenv import load_dotenv
load_dotenv()

GROQ_API_KEY = os.getenv("GROQ_API_KEY", "")
GROQ_BASE_URL = os.getenv("GROQ_BASE_URL", "https://api.groq.com/openai/v1")
GROQ_MODEL = os.getenv("GROQ_MODEL", "openai/gpt-oss-120b")
GROQ_REASONING_MODEL = os.getenv("GROQ_REASONING_MODEL", "openai/gpt-oss-120b")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY", "")
EMBEDDING_MODEL = os.getenv("EMBEDDING_MODEL", "text-embedding-3-small")
