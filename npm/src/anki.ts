const NAME = "AI Dictionary"

export async function addNote(
  Sentence: string,
  Word: string,
  PartOfSpeech: string,
  Definition: string
) {
  try {
    const decks = await invoke<string[]>("deckNames")
    if (!decks.includes(NAME)) {
      console.log("Can't find deck, skip.")
      return
    }

    const params = {
      note: {
        deckName: NAME,
        modelName: NAME,
        fields: {
          ID: new Date().toISOString(),
          Sentence,
          Word,
          PartOfSpeech,
          Definition,
        },
      },
    }
    await invoke("addNote", params)
  } catch (error) {
    console.error(error)
  }
}

async function invoke<T>(action: string, params = {}) {
  const response = await fetch("http://127.0.0.1:8765", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ action, version: 6, params }),
  })

  type body = { result: null; error: string } | { result: T; error: null }
  const body = (await response.json()) as body

  if (typeof body.error === "string") {
    throw new Error(body.error)
  }

  return body.result
}
