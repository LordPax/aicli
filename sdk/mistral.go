package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type MistralErrorMsg struct {
	Message string `json:"message"`
}

type MistralText struct {
	Sdk
	SdkText
	TextHistory
}

func NewMistralText(apiKey, model string, temp float64) (*MistralText, error) {
	history, err := NewTextHistory("mistral")
	if err != nil {
		return nil, err
	}

	sdkService := &MistralText{
		Sdk: Sdk{
			Name:   "mistral",
			ApiUrl: "https://api.mistral.ai/v1/chat/completions",
			ApiKey: apiKey,
		},
		SdkText: SdkText{
			Model: "mistral-medium",
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

func (m *MistralText) SendRequest(text string) (Message, error) {
	var textResponse OpenaiResponse

	m.AppendHistory("user", text)

	jsonBody, err := json.Marshal(OpenaiBody{
		Model:    m.Model,
		Messages: m.GetHistory(),
		Temp:     m.Temp,
	})
	if err != nil {
		return Message{}, err
	}

	resp, err := utils.PostRequest(m.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + m.ApiKey,
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
		var errorMsg MistralErrorMsg
		if err := json.Unmarshal(respBody, &errorMsg); err != nil {
			return Message{}, err
		}
		return Message{}, errors.New(errorMsg.Message)
	}

	if err := json.Unmarshal(respBody, &textResponse); err != nil {
		return Message{}, err
	}

	msg := textResponse.Choices[0].Message
	respMessage := m.AppendHistory(msg.Role, msg.GetContent())

	if err := m.SaveHistory(); err != nil {
		return Message{}, err
	}

	return respMessage, nil
}
