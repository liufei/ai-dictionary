package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/atselvan/ankiconnect"
	"github.com/fatih/color"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bold := color.New(color.Bold).SprintFunc()
	for {
		fmt.Print(bold("Sentence: "))
		sentence, _ := reader.ReadString('\n')
		sentence = strings.TrimSpace(sentence)

		fmt.Print(bold("Word: "))
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		fmt.Println()

		if sentence == "" || word == "" {
			fmt.Println(bold("Invalid input, start over."))
			fmt.Println()
			continue
		}

		fmt.Println(bold("Thinking..."))
		fmt.Println()

		meaning := ai(sentence, word)
		parts := strings.SplitN(meaning, "\n", 2)
		// partOfSpeech, definition := parts[0], parts[1]
		partOfSpeech := parts[0]
		partOfSpeech = strings.TrimSpace(partOfSpeech)
		definition := parts[1]
		definition = strings.TrimSpace(definition)

		fmt.Printf("%s%s\n", bold("Part of Speech: "), partOfSpeech)
		fmt.Printf("%s%s\n", bold("Definition: "), definition)

		anki(sentence, word, partOfSpeech, definition)
		fmt.Println()
	}
}

func ai(sentence string, word string) string {
	var baseURL option.RequestOption
	if o, ok := os.LookupEnv("OPENAI_BASE_URL"); ok {
		baseURL = option.WithBaseURL(o)
	}
	client := openai.NewClient(baseURL)

	var model string
	if m, ok := os.LookupEnv("OPENAI_MODEL"); ok {
		model = m
	}
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(`When given an English sentence and a specific word from it, follow these steps:
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
Extremely happy and excited because of something good that happened.`),
			openai.UserMessage(fmt.Sprintf(`Sentence: %s
Word: %s`, sentence, word)),
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return chatCompletion.Choices[0].Message.Content
}

func anki(sentence string, word string, partOfSpeech string, definition string) {
	const NAME = "AI Dictionary"
	client := ankiconnect.NewClient()
	decks, err := client.Decks.GetAll()
	if err != nil || !slices.Contains(*decks, NAME) {
		fmt.Println("Can't find deck, skip.")
		return
	}

	note := ankiconnect.Note{
		DeckName:  NAME,
		ModelName: NAME,
		Fields: ankiconnect.Fields{
			"ID":           time.Now().Format(time.RFC3339),
			"Sentence":     sentence,
			"Word":         word,
			"PartOfSpeech": partOfSpeech,
			"Definition":   definition,
		},
	}
	if err := client.Notes.Add(note); err != nil {
		fmt.Println("Failed to create note.")
		return
	}
}
