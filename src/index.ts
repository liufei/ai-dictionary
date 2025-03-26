#!/usr/bin/env node
import * as readline from 'node:readline/promises'
import { stdin as input, stdout as output } from 'node:process'
import { ai } from './ai.ts'
import { addNote } from './anki.ts'

const rl = readline
  .createInterface({ input, output })
  .on('SIGINT', () => rl.close())

async function run() {
  const sentence = await rl.question('Sentence: ')
  const word = await rl.question('Word: ')

  console.log()
  console.log('Thinking...')
  console.log()

  const meaning = await ai(sentence, word)
  let [partOfSpeech, definition] = meaning.split('\n')
  console.log(`Part of Speech: ${partOfSpeech}`)
  console.log(`Definition: ${definition}`)
  console.log()

  const noteId = await addNote(sentence, word, partOfSpeech, definition)
  if (noteId) console.log(`Node ID: ${noteId}`)
  console.log()

  run()
}

run()
