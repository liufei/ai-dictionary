package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
)

var (
	openaiBaseURL string
	openaiApiKey  string
	openaiModel   string

	openaiClient openai.Client
)

func init() {
	if url, ok := os.LookupEnv("OPENAI_BASE_URL"); !ok {
		log.Fatal("Can't find environment variable OPENAI_BASE_URL.")
	} else {
		openaiBaseURL = url
	}

	if key, ok := os.LookupEnv("OPENAI_API_KEY"); !ok {
		log.Fatal("Can't find environment variable OPENAI_API_KEY.")
	} else {
		openaiApiKey = key
	}

	if model, ok := os.LookupEnv("OPENAI_MODEL"); !ok {
		log.Fatal("Can't find environment variable OPENAI_MODEL.")
	} else {
		openaiModel = model
	}

	openaiClient = openai.NewClient(option.WithBaseURL(openaiBaseURL))
}

const systemMessage = `**Your Role:** You are an AI language assistant specializing in explaining words within their specific sentence context.

**Your Task:** When given an English sentence and a target word from that sentence, you must perform the following steps precisely:
1.  **Analyze Context:** Carefully examine the sentence to understand exactly how the target word is being used and what it means *in that specific situation*. This is crucial - the definition must fit the context.
2.  **Identify Part of Speech:** Determine the accurate part of speech (e.g., noun, verb, adjective, adverb) of the target word *as used in the sentence*.
3.  **Define Simply & Concisely:** Create a very brief definition (1-2 short lines maximum).
    *   Use only simple, common, everyday words, like those found in the **Oxford 3000 American English list**.
    *   The definition *must* accurately reflect the word's meaning *in the given sentence*.
    *   Avoid technical jargon or complex synonyms.

**Output Requirements:**
*   **Line 1:** The identified Part of Speech.
*   **Line 2:** The simple, context-specific definition.
*   **Strict Format:** Do not include *any* other text, explanations, greetings, or formatting beyond these two required lines.

**Example:**
Input:
sentence: She felt elated after winning the race.
word: elated

Output:
adjective
Very happy and excited because something good happened.`

func ai(sentence string, word string) (partOfSpeech, definition string) {
	chatCompletion, err := openaiClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openaiModel,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMessage),
			openai.UserMessage(fmt.Sprintf("sentence: %s\nword: %s", sentence, word)),
		},
		Temperature: param.Opt[float64]{Value: 0.1},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	meaning := strings.Split(chatCompletion.Choices[0].Message.Content, "\n")
	return meaning[0], meaning[1]
}
