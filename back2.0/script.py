import sys
import json
import google.generativeai as genai

def str_to_ascii(text):
  ascii_codes = []
  for char in text:
    ascii_codes.append(ord(char))
  return ' '.join([str(code) for code in ascii_codes])


def main():
    genai.configure(api_key="AIzaSyCPVPRR-YSDQhrRXP2qnlTnut1pbmwJ6M4")

    generation_config = {
        "temperature": 0.7,
        "top_p": 1,
        "top_k": 1,
        "max_output_tokens": 2048,
    }

    model = genai.GenerativeModel(
        model_name="gemini-1.5-pro",
        generation_config=generation_config,
    )

    chat_history = [
        {
            "role": "user",
            "parts": ["Ты учитель русского, английского и казахского языка. Исправь смысловые ошибки "],
        },
        {
            "role": "user",
            "parts": ["Выдавай только исправленние варианты. Следуй правилам языков."],
        },
    ]

    chat_session = model.start_chat(history=chat_history)

    try:
        user_sentence = sys.argv[1] if len(sys.argv) > 1 else input("Введите предложение для исправления: ")
        response = chat_session.send_message(user_sentence)
        text = str_to_ascii(response.text)
        print(text)
    except Exception as e: 
        print(f"An error occurred: {e}")


if __name__ == "__main__":
    main()
