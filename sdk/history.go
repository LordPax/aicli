package sdk

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path"

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
	AppendTextMessage(index int, text ...string) Message
	AppendImageHistory(role, fileType string, file []byte) error
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

func (t *TextHistory) AppendImageHistory(role, fileType string, file []byte) error {
	name := t.SelectedHistory

	str := base64.StdEncoding.EncodeToString(file)

	message := Message{
		Role:    role,
		Content: []IContent{NewContentImage(str, fileType)},
	}
	t.History[name] = append(t.History[name], message)

	return nil
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

	var tempHistory map[string][]json.RawMessage
	if err := json.Unmarshal(f, &tempHistory); err != nil {
		return err
	}

	t.History = make(map[string][]Message)
	for key, rawMessages := range tempHistory {
		t.History[key] = make([]Message, len(rawMessages))
		for i, rawMessage := range rawMessages {
			if err := json.Unmarshal(rawMessage, &t.History[key][i]); err != nil {
				return err
			}
		}
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

func (t *TextHistory) AppendTextMessage(index int, text ...string) Message {
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

func textContent(text ...string) []IContent {
	var content []IContent

	for _, t := range text {
		content = append(content, NewContentText(t))
	}

	return content
}
