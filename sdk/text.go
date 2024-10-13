package sdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkTextInstance ITextService

type IContent interface {
	GetValue() string
}

type ContentText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewContentText(text string) *ContentText {
	return &ContentText{
		Type: "text",
		Text: text,
	}
}

func (c *ContentText) GetValue() string {
	return c.Text
}

type ContentImage struct {
	Type   string `json:"type"`
	Source struct {
		Type      string `json:"type"`
		MediaType string `json:"media_type"`
		Data      string `json:"data"`
	} `json:"source"`
}

func NewContentImage(data, fileType string) *ContentImage {
	return &ContentImage{
		Type: "image",
		Source: struct {
			Type      string `json:"type"`
			MediaType string `json:"media_type"`
			Data      string `json:"data"`
		}{
			Type:      "base64",
			MediaType: fileType,
			Data:      data,
		},
	}
}

func (c *ContentImage) GetValue() string {
	return "image: " + c.Source.MediaType
}

type Message struct {
	Role    string     `json:"role"`
	Content []IContent `json:"content"`
}

type AuxMessage struct {
	Role    string          `json:"role"`
	Content json.RawMessage `json:"content"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var aux AuxMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	m.Role = aux.Role
	m.Content = []IContent{}

	var contentArray []json.RawMessage
	if err := json.Unmarshal(aux.Content, &contentArray); err != nil {
		return err
	}

	for _, item := range contentArray {
		content, err := unmarshalContent(item)
		if err != nil {
			return err
		}
		m.Content = append(m.Content, content)
	}

	return nil
}

func unmarshalContent(data []byte) (IContent, error) {
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}

	if _, ok := temp["text"]; ok {
		var ct ContentText
		if err := json.Unmarshal(data, &ct); err != nil {
			return nil, err
		}
		return &ct, nil
	}

	if _, ok := temp["source"]; ok {
		var ci ContentImage
		if err := json.Unmarshal(data, &ci); err != nil {
			return nil, err
		}
		return &ci, nil
	}

	return nil, fmt.Errorf("unknown content type")
}

func (m *Message) IsEmpty() bool {
	return len(m.Content) == 0
}

func (m *Message) GetContent() string {
	var text string

	if len(m.Content) == 1 {
		return m.Content[0].GetValue()
	}

	for i, c := range m.Content {
		if i != 0 {
			text += "\n---\n"
		}
		text += c.GetValue()
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
