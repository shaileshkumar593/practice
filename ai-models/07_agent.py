import json
from common.openai_client import client, DEFAULT_MODEL

CATALOG = [
    {"id": "p1", "name": "Laptop Pro", "price": 1200},
    {"id": "p2", "name": "Developer Laptop", "price": 1500},
]

def search_products(query):
    return CATALOG

def calculate_total(product_id, quantity):
    product = next(p for p in CATALOG if p["id"] == product_id)
    return {"product": product, "quantity": quantity,
            "total": product["price"] * quantity}

TOOLS = [
    {
        "type": "function", "name": "search_products",
        "description": "Search products.",
        "parameters": {
            "type": "object",
            "properties": {"query": {"type": "string"}},
            "required": ["query"], "additionalProperties": False,
        }, "strict": True,
    },
    {
        "type": "function", "name": "calculate_total",
        "description": "Calculate total price.",
        "parameters": {
            "type": "object",
            "properties": {
                "product_id": {"type": "string"},
                "quantity": {"type": "integer"},
            },
            "required": ["product_id", "quantity"],
            "additionalProperties": False,
        }, "strict": True,
    },
]

def run_agent(goal, max_steps=5):
    r = client.responses.create(
        model=DEFAULT_MODEL,
        tools=TOOLS,
        input="Achieve this goal using available tools. Never purchase anything. " + goal,
    )
    for _ in range(max_steps):
        calls = [x for x in r.output if x.type == "function_call"]
        if not calls:
            return r.output_text
        outputs = []
        for c in calls:
            args = json.loads(c.arguments)
            if c.name == "search_products":
                result = search_products(args["query"])
            elif c.name == "calculate_total":
                result = calculate_total(args["product_id"], args["quantity"])
            else:
                raise ValueError(c.name)
            outputs.append({
                "type": "function_call_output",
                "call_id": c.call_id,
                "output": json.dumps(result),
            })
        r = client.responses.create(
            model=DEFAULT_MODEL, previous_response_id=r.id, input=outputs
        )
    return "Maximum agent steps reached."

if __name__ == "__main__":
    print(run_agent("Find a developer laptop and calculate quantity 2."))


"""
    Purpose: Achieve a goal through an iterative loop of reasoning, tool use and observation.

A simple agent:

             ┌──────────────┐
             │    Goal      │
             └──────┬───────┘
                    ↓
             ┌──────────────┐
             │     LLM      │
             └──────┬───────┘
                    ↓
              Select action
                    ↓
             ┌──────────────┐
             │     Tool     │
             └──────┬───────┘
                    ↓
                 Result
                    ↓
             ┌──────────────┐
             │     LLM      │
             └──────┬───────┘
                    ↓
             Continue / Stop
    An agent typically has:

    Goal

    State/context

    Planning/reasoning

    Tool selection

    Tool execution

    Observation

    Iteration

    Termination condition

    Example:

    Goal: Find why order ORD-1001 failed.

    Agent:
    1. Get order
    2. Get payment status
    3. Get shipment status
    4. Search error logs
    5. Correlate information
    6. Explain root cause
    Interview point:
    An agent is primarily a system/application pattern, not simply a special type of neural network.

"""