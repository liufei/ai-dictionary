#!/usr/bin/env node
import * as readline from 'node:readline/promises'
import { stdin as input, stdout as output } from 'node:process'
import { ai } from './ai.ts'
import { addNote } from './anki.ts'

const rl = readline
  .createInterface({ input, output })
  .on('SIGINT', () => rl.close())

async function main() {
  const sentence = await rl.question('Sentence: ')
  console.log()

  const word = await rl.question('Word: ')
  console.log()

  console.log('Thinking...')
  console.log()

  const meaning = await ai(sentence, word)
  const [partOfSpeech, definition] = meaning.split('\n')

  console.log(`Part of Speech: ${partOfSpeech}`)
  console.log()

  console.log(`Definition: ${definition}`)
  console.log()

  await addNote(sentence, word, partOfSpeech, definition)
  console.log()

  main()
}

main()
