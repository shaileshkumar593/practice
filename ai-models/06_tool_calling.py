import json
from common.openai_client import client, DEFAULT_MODEL

def get_weather(city):
    return {"city": city, "temperature_c": 27, "condition": "rain"}

TOOLS = [{
    "type": "function",
    "name": "get_weather",
    "description": "Get weather for a city.",
    "parameters": {
        "type": "object",
        "properties": {"city": {"type": "string"}},
        "required": ["city"],
        "additionalProperties": False,
    },
    "strict": True,
}]

def ask(question):
    r = client.responses.create(model=DEFAULT_MODEL, tools=TOOLS, input=question)
    outputs = []
    for item in r.output:
        if item.type == "function_call":
            args = json.loads(item.arguments)
            result = get_weather(args["city"])
            outputs.append({
                "type": "function_call_output",
                "call_id": item.call_id,
                "output": json.dumps(result),
            })
    if not outputs:
        return r.output_text
    final = client.responses.create(
        model=DEFAULT_MODEL,
        previous_response_id=r.id,
        input=outputs,
    )
    return final.output_text

if __name__ == "__main__":
    print(ask("What is the weather in Ranchi?"))


"""

    Purpose: Allow the model to interact with external systems through controlled tools.

    The model doesn't directly access your database or API. Instead, it requests a tool call.

        User
        ↓
        LLM
        ↓
        "Call get_order_status"
        ↓
        Application
        ↓
        Order Service / DB
        ↓
        Tool Result
        ↓
        LLM
        ↓
        Final Answer
    Example tool:

    def get_order_status(order_id):
        return {
            "order_id": order_id,
            "status": "SHIPPED"
        }
    Typical tools:

    Database queries

    REST APIs

    Search

    Calculators

    CRM

    Payment APIs

    Ticketing systems

    Internal microservices

    Important security rule:

    LLM ≠ Security Boundary
    Your application must perform:

    Authentication

    Authorization

    Input validation

    Tool allowlisting

    Rate limiting

    Audit logging
"""