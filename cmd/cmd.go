package cmd

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CMD
func NewCLI() *tea.Program {
	return tea.NewProgram(initialModel(), tea.WithAltScreen())
}

// TUI
type state byte

const (
	stateInitial state = iota
	statePending
	stateFulfilled
	stateRejected
)

type focus byte

const (
	focusSentence focus = iota
	focusWord
)

type model struct {
	termWidth int

	state state
	focus focus

	sentenceInput textarea.Model
	wordInput     textarea.Model
	spinner       spinner.Model

	sentence     string
	word         string
	partOfSpeech string
	definition   string

	quitting bool
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

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

type aiMsg struct {
	sentence     string
	word         string
	partOfSpeech string
	definition   string
}

func askAI(sentence, word string) tea.Cmd {
	return func() tea.Msg {
		partOfSpeech, definition := ai(sentence, word)

		if saveAnki {
			anki(sentence, word, partOfSpeech, definition)
		}

		return aiMsg{
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

	switch m.state {
	case stateInitial, stateFulfilled, stateRejected:
		switch m.focus {
		case focusSentence:
			if keyMsg, ok := msg.(tea.KeyMsg); ok {
				if key.Matches(keyMsg, key.NewBinding(key.WithKeys("tab"))) {
					m.sentenceInput.Blur()
					m.wordInput.Focus()
					m.focus = focusWord
					return m, textarea.Blink
				}
			}
			m.sentenceInput, cmd = m.sentenceInput.Update(msg)
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		case focusWord:
			if keyMsg, ok := msg.(tea.KeyMsg); ok {
				switch {
				case key.Matches(keyMsg, key.NewBinding((key.WithKeys("tab")))):
					m.wordInput.Blur()
					m.sentenceInput.Focus()
					m.focus = focusSentence
					return m, textarea.Blink
				case key.Matches(keyMsg, key.NewBinding(key.WithKeys("enter"))):
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
						cmds = append(cmds, m.spinner.Tick, askAI(cleanedSentence, cleanedWord))
						return m, tea.Batch(cmds...)
					} else {
						m.state = stateRejected
						return m, nil
					}
				}
			}
			m.wordInput, cmd = m.wordInput.Update(msg)
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		}
	case statePending:
		switch msg := msg.(type) {
		case spinner.TickMsg:
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		case aiMsg:
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
		r.block(m.spinner.View() + " Thinking...")
	case stateFulfilled:
		r.block("Sentence:")
		r.inline(promptRender)
		r.block(styleFull.Render(m.sentence))
		r.WriteString("\n")

		r.block("Word:")
		r.inline(promptRender)
		r.block(m.word)
		r.WriteString("\n")

		r.block("Part of Speech:")
		r.inline(promptRender)
		r.block(m.partOfSpeech)
		r.WriteString("\n")

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
