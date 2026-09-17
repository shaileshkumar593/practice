from common.openai_client import client, REASONING_MODEL

def solve(problem: str) -> str:
    r = client.responses.create(
        model=REASONING_MODEL,
        reasoning={"effort": "high"},
        input=problem,
    )
    return r.output_text

if __name__ == "__main__":
    print(solve(
        "Design a payment service handling duplicate requests, provider timeouts, "
        "out-of-order webhooks and reconciliation."
    ))


"""
    Reasoning Model
        Purpose: Solve complex, multi-step problems.

        Designed to spend more computation on difficult problems.

        Useful for system design, mathematics, planning, debugging and complex decision workflows.

        Can break a problem into multiple logical steps internally.

        Usually more useful for difficult tasks than simply increasing prompt length.

        Reasoning effort may be configurable depending on the model/provider.

        Problem
        ↓
        Analyze constraints
        ↓
        Consider alternatives
        ↓
        Evaluate trade-offs
        ↓
        Produce solution


        Interview point:
            A reasoning model is still an LLM-family model, but is optimized for complex problem solving rather than only fast text generation.
"""