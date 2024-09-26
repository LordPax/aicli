package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type OpenaiResponse struct {
	Choices Choices `json:"choices"`
}

type Choices []struct {
	Index   int64          `json:"index"`
	Message ChoicesMessage `json:"message"`
}

type ChoicesMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (m *ChoicesMessage) GetContent() string {
	return m.Content
}

type OpenaiBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Temp     float64   `json:"temperature"`
}

type OpenaiText struct {
	Sdk
	SdkText
	TextHistory
}

// Initialize OpenaiText struct, inheriting from Sdk and SdkText
func NewOpenaiText(apiKey, model string, temp float64) (*OpenaiText, error) {
	history, err := NewTextHistory("openai")
	if err != nil {
		return nil, err
	}

	sdkService := &OpenaiText{
		Sdk: Sdk{
			Name:   "openai",
			ApiUrl: "https://api.openai.com/v1/chat/completions",
			ApiKey: apiKey,
		},
		SdkText: SdkText{
			Model: "gpt-4",
			Temp:  0.7,
		},
		TextHistory: *history,
	}

	if model != "" {
		sdkService.Model = model
	}

	if temp != 0 {
		sdkService.Temp = temp
	}

	if err := sdkService.LoadHistory(); err != nil {
		return nil, err
	}

	return sdkService, nil
}

func (o *OpenaiText) SendRequest(text string) (Message, error) {
	var textResponse OpenaiResponse

	o.AppendHistory("user", text)

	jsonBody, err := json.Marshal(OpenaiBody{
		Model:    o.Model,
		Messages: o.GetHistory(),
		Temp:     o.Temp,
	})
	if err != nil {
		return Message{}, err
	}

	resp, err := utils.PostRequest(o.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.ApiKey,
	})
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errorMsg ErrorMsg
		if err := json.Unmarshal(respBody, &errorMsg); err != nil {
			return Message{}, err
		}
		return Message{}, errors.New(errorMsg.Error.Message)
	}

	if err := json.Unmarshal(respBody, &textResponse); err != nil {
		return Message{}, err
	}

	msg := textResponse.Choices[0].Message
	respMessage := o.AppendHistory(msg.Role, msg.GetContent())

	if err := o.SaveHistory(); err != nil {
		return Message{}, err
	}

	return respMessage, nil
}
