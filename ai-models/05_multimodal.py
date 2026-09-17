import sys
from common.openai_client import client, DEFAULT_MODEL

def analyze(image_url, question):
    r = client.responses.create(
        model=DEFAULT_MODEL,
        input=[{"role": "user", "content": [
            {"type": "input_text", "text": question},
            {"type": "input_image", "image_url": image_url},
        ]}],
    )
    return r.output_text

if __name__ == "__main__":
    print(analyze(
        sys.argv[1],
        "Describe this image and explain any visible chart trend."
    ))


"""
    Purpose: Understand or generate multiple types of data.

    Modalities can include:

        Text
        Images
        Audio
        Video
        Documents
                ↓
        Multimodal Model
                ↓
        Understanding / Generation
    For example:

    Invoice image
        ↓
    Multimodal Model
        ↓
    Invoice number
    Amount
    Date
    Supplier
    Use cases:

    Image understanding

    OCR/document analysis

    Chart interpretation

    Voice assistants

    Video understanding

    Visual question answering

    Interview point:
        A multimodal system isn't necessarily a single model for every modality; architectures can combine specialized encoders and a shared language/reasoning component.

"""