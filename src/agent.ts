import ollama from 'ollama'

export async function agent(sentence: string, word: string) {
  const { response } = await ollama.generate({
    model: 'gemma3:12b',
    prompt: `Sentence: ${sentence}
Word: ${word}`,
    system: `When given an English sentence and a specific word from it, follow these steps:
1. Analyze the sentence's context to understand the word's role.
2. Identify the part of speech (noun, verb, adjective, etc.) accurately.
3. Define the word using simple, everyday language from the Oxford 3000 American English list.
4. Keep the definition concise (1-2 short lines, like a text message).

Output Format:
[correct part of speech]
[clear, simple explanation]  

Example:

Input:
sentence: She felt elated after winning the race.
word: elated

Output:
adjective
Extremely happy and excited because of something good that happened.`,
    stream: false,
  })
  return response
}
