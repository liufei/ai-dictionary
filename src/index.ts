#!/usr/bin/env node
import * as readline from "node:readline/promises";
import { stdin as input, stdout as output } from "node:process";
import { agent } from "./agent.ts";

const rl = readline
  .createInterface({ input, output })
  .on("SIGINT", () => rl.close());

async function run() {
  const sentence = await rl.question("Sentence: ");
  const word = await rl.question("Word: ");

  console.log();
  console.log("Thinking...");
  console.log();

  const definition = await agent(sentence, word);
  console.log(definition);
  console.log();

  run();
}

run();
