import { YankiConnect } from 'yanki-connect'

const NAME = 'AI Dictionary'
const client = new YankiConnect({ autoLaunch: process.platform === 'darwin' })

export async function addNote(
  Sentence: string,
  Word: string,
  PartOfSpeech: string,
  Definition: string
) {
  const decks = await client.deck.deckNames()
  if (!decks.includes(NAME)) {
    console.log("Can't find deck, skip.")
    return null
  }

  const note = {
    note: {
      deckName: NAME,
      fields: {
        ID: new Date().toISOString(),
        Sentence,
        Word,
        PartOfSpeech,
        Definition,
      },
      modelName: NAME,
    },
  }
  const noteId = await client.note.addNote(note)
  return noteId
}
