package sdk

import (
	"errors"
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkTextInstance ITextService

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
	var text string

	if len(m.Content) == 1 {
		return m.Content[0].Text
	}

	for i, c := range m.Content {
		if c.Type != "text" {
			continue
		}
		if i != 0 {
			text += "\n---\n"
		}
		text += c.Text
	}

	return text
}

type ErrorMsg struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type ITextService interface {
	ISdk
	ISdkText
	ITextHistory
}

type ISdkText interface {
	SendRequest(text string) (Message, error)
	SetModel(model string)
	GetModel() string
	SetTemp(temp float64)
	GetTemp() float64
}

type SdkText struct {
	Model string
	Temp  float64
}

func InitSdkText(sdk string) error {
	var err error

	l := lang.GetLocalize()
	sdkType, apiKey, model, temp, err := getConfigText(sdk)
	if err != nil {
		return err
	}

	switch sdkType {
	case "openai":
		sdkTextInstance, err = NewOpenaiText(apiKey, model, temp)
	case "claude":
		sdkTextInstance, err = NewClaudeText(apiKey, model, temp)
	case "mistral":
		sdkTextInstance, err = NewMistralText(apiKey, model, temp)
	default:
		return fmt.Errorf(l.Get("unknown-sdk"), sdkType)
	}

	if err != nil {
		return err
	}

	return nil
}

func getConfigText(sdkType string) (string, string, string, float64, error) {
	l := lang.GetLocalize()
	confText := config.CONFIG_INI.Section("text")

	if sdkType == "" {
		sdkType = confText.Key("type").String()
		if sdkType == "" {
			return "", "", "", 0, errors.New(l.Get("type-required"))
		}
	}

	apiKey := confText.Key("apiKey").String()
	if apiKey == "" {
		apiKey = confText.Key(sdkType + "-apiKey").String()
	}

	model := confText.Key("model").String()
	if model == "" {
		model = confText.Key(sdkType + "-model").String()
	}

	temp, _ := confText.Key("temp").Float64()
	if temp == 0 {
		temp, _ = confText.Key(sdkType + "-temp").Float64()
	}

	return sdkType, apiKey, model, temp, nil
}

func GetSdkText() ITextService {
	return sdkTextInstance
}

func SetSdkText(s ITextService) {
	sdkTextInstance = s
}

func (s *SdkText) SetModel(model string) {
	s.Model = model
}

func (s *SdkText) GetModel() string {
	return s.Model
}

func (s *SdkText) SetTemp(temp float64) {
	s.Temp = temp
}

func (s *SdkText) GetTemp() float64 {
	return s.Temp
}
