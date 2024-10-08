package sdk

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/utils"
)

type ITextHistory interface {
	AppendHistory(role string, text ...string) Message
	SaveHistory() error
	LoadHistory() error
	GetHistory() []Message
	SetSelectedHistory(name string)
	GetSelectedHistory() string
	ClearHistory()
	GetMessage(index int) *Message
	AppendMessage(index int, text ...string) Message
	GetHistoryNames() []string
}

type TextHistory struct {
	History         map[string][]Message
	SelectedHistory string
	HistoryFile     string
}

func NewTextHistory(sdk string) (*TextHistory, error) {
	historyFile := path.Join(config.CONFIG_DIR, sdk+"-history.json")
	log, err := utils.GetLog()
	if err != nil {
		return nil, err
	}

	if !utils.FileExist(historyFile) {
		if err := os.WriteFile(historyFile, []byte(config.HISTORY_CONTENT), 0644); err != nil {
			return nil, err
		}
		log.Logf("History file created at %s\n", historyFile)
	}

	history := &TextHistory{
		History:         make(map[string][]Message),
		SelectedHistory: "default",
		HistoryFile:     historyFile,
	}

	return history, nil
}

func (t *TextHistory) SetSelectedHistory(name string) {
	t.SelectedHistory = name
}

func (t *TextHistory) GetSelectedHistory() string {
	return t.SelectedHistory
}

func (t *TextHistory) AppendHistory(role string, text ...string) Message {
	name := t.SelectedHistory

	message := Message{
		Role:    role,
		Content: textContent(text...),
	}
	t.History[name] = append(t.History[name], message)

	return message
}

func (t *TextHistory) SaveHistory() error {
	f, err := os.Create(t.HistoryFile)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonHistory, err := json.Marshal(t.History)
	if err != nil {
		return err
	}

	if _, err := f.Write(jsonHistory); err != nil {
		return err
	}

	return nil
}

func (t *TextHistory) LoadHistory() error {
	f, err := os.ReadFile(t.HistoryFile)
	if err != nil {
		return err
	}

	if len(f) == 0 {
		return nil
	}

	if err := json.Unmarshal(f, &t.History); err != nil {
		return err
	}

	return nil
}

func (t *TextHistory) ClearHistory() {
	name := t.SelectedHistory
	t.History[name] = []Message{}
}

func (t *TextHistory) GetHistory() []Message {
	name := t.SelectedHistory
	return t.History[name]
}

func (t *TextHistory) GetMessage(index int) *Message {
	if index < 0 {
		return nil
	}
	name := t.SelectedHistory
	return &t.History[name][index]
}

func (t *TextHistory) AppendMessage(index int, text ...string) Message {
	name := t.SelectedHistory
	message := t.GetMessage(index)

	content := textContent(text...)
	message.Content = append(message.Content, content...)

	t.History[name][index] = *message

	return *message
}

func (t *TextHistory) GetHistoryNames() []string {
	var names []string

	for k := range t.History {
		names = append(names, k)
	}

	return names
}

func textContent(text ...string) []Content {
	var content []Content

	for _, t := range text {
		content = append(content, Content{
			Type: "text",
			Text: strings.TrimSpace(t),
		})
	}

	return content
}
