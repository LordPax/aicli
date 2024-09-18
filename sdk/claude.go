package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type ClaudeResponse struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type ClaudeText struct {
	Sdk
	SdkText
}

// Initialize ClaudeText struct, inheriting from Sdk and SdkText
func NewClaudeText(apiKey, model string, temp float64) (*ClaudeText, error) {
	sdkService := &ClaudeText{
		Sdk: Sdk{
			ApiUrl: "https://api.anthropic.com/v1/messages",
			ApiKey: apiKey,
			Model:  "claude-3-5-sonnet-20240620",
		},
		SdkText: SdkText{
			History:         make(map[string][]Message),
			SelectedHistory: "default",
			Temp:            0.7,
		},
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

func (c *ClaudeText) SendRequest(text string) (Message, error) {
	var textResponse ClaudeResponse

	c.AppendHistory("user", text)

	jsonBody, err := json.Marshal(TextBody{
		Model: c.Model,
		// MaxTokens: 1024,
		Messages: c.GetHistory(),
	})
	if err != nil {
		return Message{}, err
	}

	resp, err := utils.PostRequest(c.ApiUrl, jsonBody, map[string]string{
		"Content-Type":      "application/json",
		"anthropic-version": "2023-06-01",
		"x-api-key":         c.ApiKey,
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

	respMessage := c.AppendHistory(textResponse.Role, textResponse.Content[0].Text)

	if err := c.SaveHistory(); err != nil {
		return Message{}, err
	}

	return respMessage, nil
}
