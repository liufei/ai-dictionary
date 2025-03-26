# AI Dictionary

> Ask AI to explain the word in sentence and save definition to Anki.

## Requirements

- Node.js v22.6.0+
- Anki Desktop

- AnkiConnect
- deepseek-r1 API Key

## Usage

Open Anki, create deck `AI Dictionary`.

<img width="768" alt="Screenshot 2025-03-26 at 14 37 54" src="https://github.com/user-attachments/assets/7a6fbe99-5dfd-4303-ab04-cc4fcffe194a" />

Click `Tools` -> `Manage Note Types` from Anki menu, add note type `AI Dictionary`.

<img width="768" alt="Screenshot 2025-03-26 at 14 38 43" src="https://github.com/user-attachments/assets/ceaa06f1-85cc-4a16-8812-dc05a8e4f27e" />

Add card template.

<img width="768" alt="Screenshot 2025-03-26 at 14 39 09" src="https://github.com/user-attachments/assets/76da63ae-ce04-4252-abd1-5c4994d288ed" />

Open terminal, set environment variables.

```sh
export OPENAI_API_KEY=sk-***
export OPENAI_BASE_URL=https://***
```

Paste sentence and word into command line, press `Ctrl+C` to exit.

```sh
npx ai-dictionary
```

<img width="1512" alt="Screenshot 2025-03-26 at 11 58 32" src="https://github.com/user-attachments/assets/c2c1d9e9-d543-45c4-a7e9-853045887db0" />
<img width="768" alt="Screenshot 2025-03-26 at 12 03 43" src="https://github.com/user-attachments/assets/16d1f6ac-fbf8-46b1-b7d8-402daffa062e" />
