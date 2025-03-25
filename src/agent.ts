async function fetchDefinitions(word: string) {
  const response = await fetch(
    `https://api.dictionaryapi.dev/api/v2/entries/en/${word}`
  );
  const definitions = await response.json();
  return definitions;
}

async function matchDefinition(
  sentence: string,
  word: string,
  definitions: string
) {
  const response = await fetch("http://localhost:11434/api/generate", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      model: "gemma3:12b",

      system: `You will receive:
  1. An English sentence
  2. Target word to analyze
  3. JSON dictionary data containing multiple parts of speech and definitions
  
  Task:
  1. Analyze the sentence context to determine:
     - The grammatical role of the target word
     - Its most likely part of speech
  2. From the JSON data:
     - First identify the correct partOfSpeech category
     - Then select the most context-appropriate definition within that category
  3. Explain your reasoning by:
     - Highlighting contextual clues from the sentence
     - Eliminating irrelevant parts of speech
     - Comparing with similar definitions
     - Mentioning any example matches
  
  Output Format:
  {partOfSpeech}: {selected definition} (exact wording from JSON)
  
  Example Input:
  Sentence: "Hello? How may I help you?"
  Word: "hello"
  JSON: [{"word":"hello","meanings":[{"partOfSpeech":"noun","definitions":["\"Hello!\" or an equivalent greeting."]},{"partOfSpeech":"verb","definitions":["To greet with \"hello\"."]},{"partOfSpeech":"interjection","definitions":["A greeting (salutation) said when meeting someone or acknowledging someone’s arrival or presence.","A greeting used when answering the telephone.","A call for response if it is not clear if anyone is present or listening, or if a telephone conversation may have been disconnected.","Used sarcastically to imply that the person addressed or referred to has done something the speaker or writer considers to be foolish.","An expression of puzzlement or discovery."]}]}]
  
  Example Output:
  interjection: A greeting used when answering the telephone.
  
  Attention:
  The output shouldn't contain any reasoning process.`,

      prompt: `Sentence: ${sentence}
  Word: ${word}
  JSON: ${definitions}`,

      stream: false,
    }),
  });
  const payload = await response.json();
  return payload.response;
}

export async function agent(sentence: string, word: string) {
  const definitions = await fetchDefinitions(word);
  const definition = await matchDefinition(
    sentence,
    word,
    reduceToken(definitions)
  );
  return definition;
}

interface IDefinition {
  word: string;
  meanings: {
    partOfSpeech: string;
    definitions: {
      definition: string;
    }[];
  }[];
}

function reduceToken(definitions: IDefinition[]) {
  const reduced = definitions.map(({ word, meanings }) => ({
    word,
    meanings: meanings.map(({ partOfSpeech, definitions }) => ({
      partOfSpeech,
      definitions: definitions.map(({ definition }) => definition),
    })),
  }));
  return JSON.stringify(reduced);
}

// agent("Hello! What’s going on here?", "hello").then(console.log);
// agent("Hello? How may I help you?", "hello").then(console.log);
// agent("Hello? Is anyone there?", "hello").then(console.log);
