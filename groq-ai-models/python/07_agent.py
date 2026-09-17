import json
from common.clients import groq

tools=[{"type":"function","function":{
    "name":"search_docs","description":"Search internal knowledge.",
    "parameters":{"type":"object","properties":{"query":{"type":"string"}},
                  "required":["query"]}
}}]

def search_docs(q):
    docs=[
        "Refunds can be initiated only for captured payments.",
        "Refund requests require an idempotency key.",
        "Webhook events must be deduplicated using the provider event ID."
    ]
    return [d for d in docs if any(w.lower() in d.lower() for w in q.split())]

def run_agent(goal,max_steps=5):
    messages=[{"role":"system","content":"Use tools when needed. Never invent tool results."},
              {"role":"user","content":goal}]
    for _ in range(max_steps):
        r=groq.chat.completions.create(model="openai/gpt-oss-120b",
                                       messages=messages,tools=tools,tool_choice="auto")
        msg=r.choices[0].message
        messages.append(msg)
        if not msg.tool_calls:
            return msg.content
        for c in msg.tool_calls:
            if c.function.name=="search_docs":
                args=json.loads(c.function.arguments)
                messages.append({"role":"tool","tool_call_id":c.id,
                                 "content":json.dumps(search_docs(args["query"]))})
    return "Maximum agent steps reached."

print(run_agent("Find the internal rules for safely processing payment refunds."))
