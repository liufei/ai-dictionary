import { useRef, useState } from 'react'
import { createRoot } from 'react-dom/client'

createRoot(document.body).render(<App />)

function App() {
  const [loading, setLoading] = useState(false)
  const [meaning, setMeaning] = useState(null)
  const formRef = useRef<HTMLFormElement>(null)

  async function handleSubmit(event: React.SyntheticEvent) {
    event.preventDefault()
    const target = event.target as typeof event.target & {
      sentence: { value: string }
      word: { value: string }
    }
    const sentence = target.sentence.value
    const word = target.word.value
    setLoading(true)
    const meaning: Meaning = await window.ai.lookup(sentence, word)
    setLoading(false)
    setMeaning(meaning)
    formRef.current?.reset()
  }

  return (
    <main>
      <form className="pure-form" onSubmit={handleSubmit} ref={formRef}>
        <fieldset className="pure-group">
          <textarea
            name="sentence"
            placeholder="Sentence"
            className="pure-input-1"
            rows={4}
            required
          ></textarea>
          <input
            name="word"
            placeholder="Word"
            className="pure-input-1"
            required
          />
        </fieldset>
        <button
          type="submit"
          className="pure-button pure-input-1"
          disabled={loading}
        >
          Ask AI to lookup dictionary
        </button>
      </form>
      {loading ? <p>Thinking...</p> : <Dictionary data={meaning} />}
    </main>
  )
}

function Dictionary({ data }: { data: Meaning }) {
  if (!data) return null

  return (
    <dl>
      <dt>Sentence</dt>
      <dd>{data.sentence}</dd>
      <dt>Word</dt>
      <dd>{data.word}</dd>
      <dt>Part of Speech</dt>
      <dd>{data.partOfSpeech}</dd>
      <dt>Definition</dt>
      <dd>{data.definition}</dd>
    </dl>
  )
}
