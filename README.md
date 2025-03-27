# AI Dictionary

> Ask AI to explain the word in sentence and save definition to Anki.

## Ask AI to explain the word in sentence

### Requirements

- Node.js
- LLM API Key

### Usage

1. Open terminal, set environment variables.

```sh
# Online: OpenAI, Deepseek, etc.
export OPENAI_BASE_URL=https://***
export OPENAI_API_KEY=sk-***
export OPENAI_MODEL=deepseek-r1

# Offline: Ollama
export OPENAI_BASE_URL=http://localhost:11434/v1
export OPENAI_API_KEY=ollama
export OPENAI_MODEL=gemma3:27b
```

2. Paste sentence and word into command line, press `Ctrl+C` to exit.

```sh
npx ai-dictionary
```

### Example

<img width="1512" alt="Screenshot 2025-03-27 at 14 02 34" src="https://github.com/user-attachments/assets/f52ec69f-8aa5-467a-9326-9ab403278a5d" />

## Save definition to Anki

### Requirements

- Anki Desktop
- AnkiConnect

### Usage

1. Open Anki, create deck `AI Dictionary`.

<img width="768" alt="Screenshot 2025-03-26 at 14 37 54" src="https://github.com/user-attachments/assets/7a6fbe99-5dfd-4303-ab04-cc4fcffe194a" />

2. Click `Tools` -> `Manage Note Types` from Anki menu, add note type `AI Dictionary`.

<img width="768" alt="Screenshot 2025-03-26 at 14 38 43" src="https://github.com/user-attachments/assets/ceaa06f1-85cc-4a16-8812-dc05a8e4f27e" />

3. Add card template.

<img width="768" alt="Screenshot 2025-03-26 at 14 39 09" src="https://github.com/user-attachments/assets/76da63ae-ce04-4252-abd1-5c4994d288ed" />

### Example

<img width="768" alt="Screenshot 2025-03-26 at 12 03 43" src="https://github.com/user-attachments/assets/16d1f6ac-fbf8-46b1-b7d8-402daffa062e" />
