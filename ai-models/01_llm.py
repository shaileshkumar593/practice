from common.openai_client import client, DEFAULT_MODEL

def generate(prompt: str) -> str:
    r = client.responses.create(model=DEFAULT_MODEL, input=prompt)
    return r.output_text

if __name__ == "__main__":
    print(generate("Explain PostgreSQL indexing in five bullet points."))



"""
    1. LLM — Large Language Model
        Purpose: Generate and understand human language.

        Trained on large amounts of text.

        Performs tasks such as text generation, summarization, translation, classification, Q&A and code generation.

        Works primarily with tokens.

        Examples: GPT, Llama, Claude, Gemini.

        Key concepts: tokens, context window, temperature, prompting, hallucination.

        In an application, the LLM is usually responsible for producing the final natural-language response.
"""