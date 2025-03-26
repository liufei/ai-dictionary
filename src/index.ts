#!/usr/bin/env node
import * as readline from 'node:readline/promises'
import { stdin as input, stdout as output } from 'node:process'
import { execSync } from 'node:child_process'
import { ai } from './ai.ts'
import { addNote } from './anki.ts'

const rl = readline
  .createInterface({ input, output })
  .on('SIGINT', () => rl.close())

async function run() {
  const Sentence = await rl.question('Sentence: ')
  const Word = await rl.question('Word: ')

  console.log()
  console.log('Thinking...')
  console.log()

  const meaning = await ai(Sentence, Word)
  execSync(`say "${Word}"`)
  console.log(meaning)
  console.log()

  const [PartOfSpeech, Definition] = meaning.split('\n')
  const noteId = await addNote(Sentence, Word, PartOfSpeech, Definition)
  console.log(`Node ID: ${noteId}`)
  console.log()

  run()
}

run()
