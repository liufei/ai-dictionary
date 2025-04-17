package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	stateInitial state = iota
	statePending
	stateFulfilled
	stateRejected
)

type focus int

const (
	focusSentence focus = iota
	focusWord
)

type model struct {
	termWidth     int
	state         state
	focus         focus
	sentenceInput textarea.Model
	wordInput     textarea.Model
	spinner       spinner.Model
	sentence      string
	word          string
	partOfSpeech  string
	definition    string
	quitting      bool
}

func initialModel() model {
	m := model{termWidth: 80}

	m.sentenceInput = textarea.New()
	m.sentenceInput.SetWidth(m.termWidth)
	m.sentenceInput.SetHeight(5)
	m.sentenceInput.Focus()
	m.sentenceInput.ShowLineNumbers = false
	m.sentenceInput.Prompt = "> "
	m.sentenceInput.Placeholder = "Paste the sentence here, and then press the Tab key"
	m.sentenceInput.KeyMap.InsertNewline.SetEnabled(true)

	m.wordInput = textarea.New()
	m.wordInput.SetWidth(m.termWidth)
	m.wordInput.SetHeight(1)
	m.wordInput.ShowLineNumbers = false
	m.wordInput.Prompt = "> "
	m.wordInput.Placeholder = "Paste the word here, and then press the Enter key"
	m.wordInput.KeyMap.InsertNewline.SetEnabled(false)

	m.spinner = spinner.New(spinner.WithSpinner(spinner.Dot))

	return m
}

type processCompleteMsg struct {
	sentence     string
	word         string
	partOfSpeech string
	definition   string
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func runAIAndNote(sentence, word string) tea.Cmd {
	return func() tea.Msg {
		partOfSpeech, definition := ai(sentence, word)

		if saveAnki {
			anki(sentence, word, partOfSpeech, definition)
		}

		return processCompleteMsg{
			sentence:     sentence,
			word:         word,
			partOfSpeech: partOfSpeech,
			definition:   definition,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.sentenceInput.SetWidth(m.termWidth)
		m.wordInput.SetWidth(m.termWidth)
		return m, nil
	case tea.KeyMsg:
		if key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c", "esc"))) {
			m.quitting = true
			return m, tea.Quit
		}
	}

	if m.state != statePending {
		switch m.focus {
		case focusSentence:
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch {
				case key.Matches(msg, key.NewBinding(key.WithKeys("tab"))):
					m.sentenceInput.Blur()
					m.wordInput.Focus()
					m.focus = focusWord
					return m, textarea.Blink
				default:
					m.sentenceInput, cmd = m.sentenceInput.Update(msg)
					cmds = append(cmds, cmd)
					return m, tea.Batch(cmds...)
				}
			}
		case focusWord:
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch {
				case key.Matches(msg, key.NewBinding((key.WithKeys("tab")))):
					m.wordInput.Blur()
					m.sentenceInput.Focus()
					m.focus = focusSentence
					return m, textarea.Blink
				case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
					sentence := m.sentenceInput.Value()
					cleanedSentence := strings.TrimSpace(strings.ReplaceAll(sentence, "\n", " "))
					word := m.wordInput.Value()
					cleanedWord := strings.TrimSpace(strings.ReplaceAll(word, "\n", " "))
					if cleanedSentence != "" && cleanedWord != "" {
						m.state = statePending
						m.sentence = ""
						m.word = ""
						m.partOfSpeech = ""
						m.definition = ""
						cmds = append(cmds, m.spinner.Tick)
						cmds = append(cmds, runAIAndNote(cleanedSentence, cleanedWord))
						return m, tea.Batch(cmds...)
					} else {
						m.state = stateRejected
						return m, nil
					}
				default:
					m.wordInput, cmd = m.wordInput.Update(msg)
					cmds = append(cmds, cmd)
					return m, tea.Batch(cmds...)
				}
			}
		}
	} else {
		switch msg := msg.(type) {
		case spinner.TickMsg:
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		case processCompleteMsg:
			m.sentenceInput.SetValue("")
			m.sentenceInput.Focus()
			m.wordInput.SetValue("")
			m.wordInput.Blur()
			m.state = stateFulfilled
			m.focus = focusSentence
			m.sentence = msg.sentence
			m.word = msg.word
			m.partOfSpeech = msg.partOfSpeech
			m.definition = msg.definition
			return m, textarea.Blink
		}
	}

	return m, nil
}

type render struct {
	strings.Builder
}

func (r *render) inline(s string) {
	r.WriteString(s)
}

func (r *render) block(s string) {
	r.WriteString(s)
	r.WriteString("\n")
}

func (m model) View() string {
	var r render

	if m.quitting {
		r.block("Exiting...")
		return r.String()
	}

	styleFull := lipgloss.NewStyle().Width(m.termWidth)
	styleCenter := styleFull.Align(lipgloss.Center)

	r.block(styleCenter.Render("AI Dictionary"))

	r.block(styleCenter.Render("https://github.com/liufei/ai-dictionary"))

	dividerRender := strings.Repeat("-", m.termWidth)

	r.block(dividerRender)

	r.block(styleCenter.Render("Tab: navigate | Enter (in Word): submit | Esc: quit"))

	r.block(dividerRender)

	r.block("Sentence:")
	r.block(m.sentenceInput.View())

	r.block("Word:")
	r.block(m.wordInput.View())

	if m.state != stateInitial {
		r.block(dividerRender)
	}

	promptRender := lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("> ")

	switch m.state {
	case statePending:
		r.block(m.spinner.View() + " Processing...")
	case stateFulfilled:
		r.block("Sentence:")
		r.inline(promptRender)
		r.block(styleFull.Render(m.sentence))
		r.block("Word:")
		r.inline(promptRender)
		r.block(m.word)
		r.block("Part of Speech:")
		r.inline(promptRender)
		r.block(m.partOfSpeech)
		r.block("Definition:")
		r.inline(promptRender)
		r.block(styleFull.Render(m.definition))
	case stateRejected:
		r.block("Error:")
		r.inline(promptRender)
		r.block("Sentence and Word are required.")
	}

	return r.String()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
