package sdk

import (
	"fmt"
)

type OpenaiText struct {
	Sdk
	SdkText
}

// Initialize OpenaiText struct, inheriting from Sdk and SdkText
func NewOpenaiText(apiKey string, model string) *OpenaiText {
	return &OpenaiText{
		Sdk: Sdk{
			apiUrl: "https://api.openai.com/v1/chat/completions",
			ApiKey: apiKey,
			Model:  model,
		},
		SdkText: SdkText{
			History: []Message{},
			Temp:    0.7,
		},
	}
}

func (o *OpenaiText) SendRequest(text string) (string, error) {
	fmt.Println("Sending request to OpenAI", text)

	// http.Post(o.apiUrl, "application/json", nil)
	return text, nil
}
