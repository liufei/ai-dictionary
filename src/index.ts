#!/usr/bin/env node
import * as readline from 'node:readline/promises'
import { stdin as input, stdout as output } from 'node:process'
import colors from 'yoctocolors'
import { ai } from './ai.ts'
import { addNote } from './anki.ts'

const rl = readline
  .createInterface({ input, output })
  .on('SIGINT', () => rl.close())

async function main() {
  const sentence = await rl.question(colors.bold('Sentence: '))
  const word = await rl.question(colors.bold('Word: '))
  console.log()

  if (!sentence || !word) {
    console.log(colors.bold('Invalid input, start over.'))
    console.log()
    return main()
  }
  console.log(colors.bold('Thinking...'))
  console.log()

  const meaning = await ai(sentence, word)
  const [partOfSpeech, definition] = meaning.split('\n')
  console.log(`${colors.bold('Part of Speech:')} ${partOfSpeech}`)
  console.log(`${colors.bold('Definition:')} ${definition}`)
  await addNote(sentence, word, partOfSpeech, definition)
  console.log()

  return main()
}

main()
