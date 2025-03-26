import { YankiConnect } from 'yanki-connect'

const client = new YankiConnect({ autoLaunch: process.platform === 'darwin' })

export async function addNote(
  Sentence: string,
  Word: string,
  PartOfSpeech: string,
  Definition: string
) {
  const note = {
    note: {
      deckName: 'AI Dictionary',
      fields: {
        ID: new Date().toISOString(),
        Sentence,
        Word,
        PartOfSpeech,
        Definition,
      },
      modelName: 'AI Dictionary',
    },
  }
  const noteId = await client.note.addNote(note)
  return noteId
}
