import json
from common.clients import groq

tools=[{"type":"function","function":{
    "name":"get_order_status",
    "description":"Get the current status of an order.",
    "parameters":{"type":"object","properties":{"order_id":{"type":"string"}},
                  "required":["order_id"]}
}}]

def get_order_status(order_id):
    return {"order_id":order_id,"status":"SHIPPED"}

messages=[{"role":"user","content":"What is the status of order ORD-1001?"}]
r=groq.chat.completions.create(model="openai/gpt-oss-120b",
                               messages=messages,tools=tools,tool_choice="auto")
msg=r.choices[0].message
messages.append(msg)

for call in msg.tool_calls or []:
    if call.function.name=="get_order_status":
        args=json.loads(call.function.arguments)
        messages.append({"role":"tool","tool_call_id":call.id,
                         "content":json.dumps(get_order_status(args["order_id"]))})

r=groq.chat.completions.create(model="openai/gpt-oss-120b",messages=messages,tools=tools)
print(r.choices[0].message.content)
