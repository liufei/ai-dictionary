#!/usr/bin/env node
import * as readline from "node:readline/promises"
import { stdin as input, stdout as output } from "node:process"
import { styleText } from "node:util"
import { ai } from "./ai.ts"
import { addNote } from "./anki.ts"

const rl = readline
  .createInterface({ input, output })
  .on("SIGINT", () => rl.close())

const bold = styleText.bind(this, "bold")

async function main() {
  const sentence = await rl.question(bold("Sentence: "))
  const word = await rl.question(bold("Word: "))
  console.log()

  if (!sentence || !word) {
    console.log(bold("Invalid input, start over."))
    console.log()
    return main()
  }
  console.log(bold("Thinking..."))
  console.log()

  const meaning = await ai(sentence, word)
  let [partOfSpeech, definition] = meaning.split("\n")
  partOfSpeech = partOfSpeech.trim()
  definition = definition.trim()
  console.log(`${bold("Part of Speech:")} ${partOfSpeech}`)
  console.log(`${bold("Definition:")} ${definition}`)
  await addNote(sentence, word, partOfSpeech, definition)
  console.log()

  return main()
}

main()
