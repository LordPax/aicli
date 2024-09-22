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

func (c *ClaudeResponse) GetContent() string {
	return c.Content[0].Text
}

type ClaudeBody struct {
	Model     string    `json:"model"`
	MaxTokens int64     `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type ClaudeText struct {
	Sdk
	SdkText
	TextHistory
}

// Initialize ClaudeText struct, inheriting from Sdk and SdkText
func NewClaudeText(apiKey, model string, temp float64) (*ClaudeText, error) {
	history, err := NewTextHistory("claude")
	if err != nil {
		return nil, err
	}

	sdkService := &ClaudeText{
		Sdk: Sdk{
			Name:   "claude",
			ApiUrl: "https://api.anthropic.com/v1/messages",
			ApiKey: apiKey,
		},
		SdkText: SdkText{
			Model: "claude-3-5-sonnet-20240620",
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

func (c *ClaudeText) SendRequest(text string) (Message, error) {
	var textResponse ClaudeResponse

	idLastMsg := len(c.GetHistory()) - 1
	lastMessage := c.GetMessage(idLastMsg)

	if lastMessage != nil && lastMessage.Role == "user" {
		c.AppendMessage(idLastMsg, text)
	} else {
		c.AppendHistory("user", text)
	}

	jsonBody, err := json.Marshal(ClaudeBody{
		Model:     c.Model,
		MaxTokens: 1024,
		Messages:  c.GetHistory(),
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

	respMessage := c.AppendHistory(textResponse.Role, textResponse.GetContent())

	if err := c.SaveHistory(); err != nil {
		return Message{}, err
	}

	return respMessage, nil
}

func (c *ClaudeText) AppendHistory(role string, text ...string) Message {
	var content []Content
	name := c.SelectedHistory

	if role == "system" {
		role = "user"
	}

	for _, t := range text {
		content = append(content, Content{
			Type: "text",
			Text: t,
		})
	}

	message := Message{
		Role:    role,
		Content: content,
	}
	c.History[name] = append(c.History[name], message)

	return message
}
