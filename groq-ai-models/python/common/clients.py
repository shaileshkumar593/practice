from openai import OpenAI
from .config import *

if not GROQ_API_KEY:
    raise RuntimeError("GROQ_API_KEY is missing")

groq = OpenAI(api_key=GROQ_API_KEY, base_url=GROQ_BASE_URL)
openai_client = OpenAI(api_key=OPENAI_API_KEY) if OPENAI_API_KEY else None

def chat(prompt, system="", model=GROQ_MODEL):
    messages = ([{"role":"system","content":system}] if system else [])
    messages.append({"role":"user","content":prompt})
    r = groq.chat.completions.create(model=model, messages=messages, temperature=0.2)
    return r.choices[0].message.content or ""

def reason(prompt):
    r = groq.chat.completions.create(
        model=GROQ_REASONING_MODEL,
        messages=[{"role":"user","content":prompt}],
        temperature=0.1,
        extra_body={"reasoning_effort":"high"},
    )
    return r.choices[0].message.content or ""

def embed(texts):
    if not openai_client:
        raise RuntimeError("OPENAI_API_KEY required for this embedding example")
    r = openai_client.embeddings.create(model=EMBEDDING_MODEL, input=texts)
    return [x.embedding for x in r.data]
