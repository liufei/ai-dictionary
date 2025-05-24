interface Window {
  ai: {
    lookup: (sentence: string, word: string) => Promise<Meaning>
  }
}

interface Meaning {
  sentence: string
  word: string
  partOfSpeech: string
  definition: string
}
