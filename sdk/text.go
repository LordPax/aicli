package sdk

import (
	"encoding/json"
	"os"

	"github.com/LordPax/aicli/config"
)

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
	// Source struct {
	// 	Type      string `json:"type"`
	// 	MediaType string `json:"media_type"`
	// 	Data      string `json:"data"`
	// } `json:"source"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

func (m *Message) GetContent() string {
	return m.Content[0].Text
}

type ErrorMsg struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type ITextService interface {
	ISdk
	ISdkText
}

type ISdkText interface {
	SetTemp(temp float64)
	GetTemp() float64
	AppendHistory(role string, text ...string) Message
	SaveHistory() error
	LoadHistory() error
	GetHistory() []Message
	SetSelectedHistory(name string)
	GetSelectedHistory() string
	ClearHistory()
}

type SdkText struct {
	History         map[string][]Message
	SelectedHistory string
	Temp            float64
}

func (s *SdkText) SetSelectedHistory(name string) {
	s.SelectedHistory = name
}

func (s *SdkText) GetSelectedHistory() string {
	return s.SelectedHistory
}

func (s *SdkText) SetTemp(temp float64) {
	s.Temp = temp
}

func (s *SdkText) GetTemp() float64 {
	return s.Temp
}

func (s *SdkText) AppendHistory(role string, text ...string) Message {
	var content []Content

	for _, t := range text {
		content = append(content, Content{
			Type: "text",
			Text: t,
		})
	}

	name := s.SelectedHistory
	message := Message{
		Role:    role,
		Content: content,
	}
	s.History[name] = append(s.History[name], message)

	return message
}

func (s *SdkText) SaveHistory() error {
	f, err := os.Create(config.HISTORY_FILE)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonHistory, err := json.Marshal(s.History)
	if err != nil {
		return err
	}

	if _, err := f.Write(jsonHistory); err != nil {
		return err
	}

	return nil
}

func (s *SdkText) LoadHistory() error {
	f, err := os.ReadFile(config.HISTORY_FILE)
	if err != nil {
		return err
	}

	if len(f) == 0 {
		return nil
	}

	if err := json.Unmarshal(f, &s.History); err != nil {
		return err
	}

	return nil
}

func (s *SdkText) ClearHistory() {
	name := s.SelectedHistory
	s.History[name] = []Message{}
}

func (s *SdkText) GetHistory() []Message {
	name := s.SelectedHistory
	return s.History[name]
}
