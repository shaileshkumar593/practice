import os
from dotenv import load_dotenv
from openai import OpenAI

load_dotenv()
client = OpenAI(api_key=os.environ["OPENAI_API_KEY"])
DEFAULT_MODEL = os.getenv("OPENAI_MODEL", "gpt-5.6-luna")
REASONING_MODEL = os.getenv("OPENAI_REASONING_MODEL", "gpt-5.6-sol")
EMBEDDING_MODEL = os.getenv("OPENAI_EMBEDDING_MODEL", "text-embedding-3-small")
