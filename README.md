# AI Dictionary

> Ask AI to explain the word in a sentence and save the definition to Anki.

## Ask AI to explain the word in a sentence

### Requirements

- Node.js >=20.12.0
- LLM API Key

### Usage

#### Browser

1. Open [https://liufei.github.io/ai-dictionary/](https://liufei.github.io/ai-dictionary/).
2. Set up the configuration for `OPENAI_BASE_URL`, `OPENAI_API_KEY`, and `OPENAI_MODEL`, which will be saved in your browser's `localStorage`.
3. Paste sentence and word into form.

#### Terminal

1. Open terminal, set environment variables.
   - OpenAI, Deepseek, etc.
     ```sh
     export OPENAI_BASE_URL=https://api.deepseek.com/v1; export OPENAI_API_KEY=sk-***; export OPENAI_MODEL=deepseek-r1
     ```
   - Ollama
     ```sh
     export OPENAI_BASE_URL=http://localhost:11434/v1; export OPENAI_API_KEY=ollama; export OPENAI_MODEL=gemma3:27b
     ```
2. Run command `npx ai-dictionary`.
3. Paste sentence and word into command line.
   - press `Enter` to ask AI
   - press `Ctrl+L` to clear console
   - press `Ctrl+C` to exit program

### Example

<img width="1512" alt="Screenshot 2025-03-27 at 14 02 34" src="https://github.com/user-attachments/assets/f52ec69f-8aa5-467a-9326-9ab403278a5d" />

## Save the definition to Anki

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

4. Click `Tools` -> `Add-ons` from Anki menu, select `AnkiConnect` and click `Config`, add `"https://liufei.github.io"` to the list of `"webCorsOriginList"`.

<img width="768" alt="Screenshot 2025-04-11 at 16 44 17" src="https://github.com/user-attachments/assets/69b5863a-7395-4c05-ae0b-ee3a80eb0764" />

<img width="768" alt="Screenshot 2025-04-11 at 16 44 20" src="https://github.com/user-attachments/assets/0f005af5-3239-4b85-9100-b52fb7821ef4" />

### Example

<img width="768" alt="Screenshot 2025-03-26 at 12 03 43" src="https://github.com/user-attachments/assets/16d1f6ac-fbf8-46b1-b7d8-402daffa062e" />
