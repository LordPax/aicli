package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type OpenaiText struct {
	Sdk
	SdkText
}

// Initialize OpenaiText struct, inheriting from Sdk and SdkText
func NewOpenaiText(apiKey, model string, temp float64) (*OpenaiText, error) {
	sdkService := &OpenaiText{
		Sdk: Sdk{
			ApiUrl: "https://api.openai.com/v1/chat/completions",
			ApiKey: apiKey,
			Model:  "gpt-4",
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

func (o *OpenaiText) SendRequest(text string) (Message, error) {
	var textResponse TextResponse

	message := Message{
		Role:    "user",
		Content: text,
	}

	o.AppendHistory(message)

	body := TextBody{
		Model:    o.Model,
		Messages: o.GetHistory(),
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return Message{}, err
	}

	resp, err := utils.PostRequest(o.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", o.ApiKey),
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

	respMessage := textResponse.Choices[0].Message
	o.AppendHistory(respMessage)

	if err := o.SaveHistory(); err != nil {
		return Message{}, err
	}

	return respMessage, nil
}
