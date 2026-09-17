import os, base64
from common.clients import groq

path = os.getenv("IMAGE_PATH")
if not path:
    raise SystemExit("Set IMAGE_PATH=/path/to/image")

with open(path,"rb") as f:
    b64 = base64.b64encode(f.read()).decode()

r = groq.chat.completions.create(
    model=os.getenv("GROQ_MODEL","openai/gpt-oss-120b"),
    messages=[{"role":"user","content":[
        {"type":"text","text":"Describe this image and extract useful text if present."},
        {"type":"image_url","image_url":{"url":f"data:image/png;base64,{b64}"}}
    ]}]
)
print(r.choices[0].message.content)
