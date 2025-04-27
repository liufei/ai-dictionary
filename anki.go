package dictionary

import (
	"fmt"
	"slices"
	"time"

	"github.com/atselvan/ankiconnect"
)

var (
	SaveAnki   = true
	ankiClient = ankiconnect.NewClient()
)

const NAME = "AI Dictionary"

func init() {
	decks, err := ankiClient.Decks.GetAll()
	if err != nil || !slices.Contains(*decks, NAME) {
		fmt.Println("Can't find Anki Deck, skip.")
		SaveAnki = false
	}
}

func Anki(sentence, word, partOfSpeech, definition string) {
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
	if err := ankiClient.Notes.Add(note); err != nil {
		fmt.Printf("Failed to create Anki note: %v\n", err)
		return
	}
}
