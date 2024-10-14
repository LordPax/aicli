package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type OpenaiImageBody struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

type OpenaiImageResponse struct {
	Images []struct {
		Url     string `json:"url"`
		B64Json string `json:"b64_json"`
	} `json:"images"`
}

type OpenaiImage struct {
	Sdk
	SdkImage
}

func NewOpenaiImage(apiKey string) (*OpenaiImage, error) {
	return &OpenaiImage{
		Sdk: Sdk{
			Name:   "openai",
			ApiUrl: "https://api.openai.com/v1/images/generations",
			ApiKey: apiKey,
			Inerte: false,
		},
		SdkImage: SdkImage{
			Model:   "dall-e-3",
			Size:    "1024x1024",
			ImageNb: 1,
		},
	}, nil
}

func (o *OpenaiImage) SendRequest(prompt string) (OpenaiImageResponse, error) {
	var openaiResponse OpenaiImageResponse
	format := "url"

	if o.GetOutput() != "" {
		format = "b4_json"
	}

	jsonBody, err := json.Marshal(OpenaiImageBody{
		Model:          o.GetModel(),
		Prompt:         prompt,
		N:              o.GetImageNb(),
		Size:           o.GetSize(),
		ResponseFormat: format,
	})
	if err != nil {
		return OpenaiImageResponse{}, err
	}

	resp, err := utils.PostRequest(o.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.ApiKey,
	})
	if err != nil {
		return OpenaiImageResponse{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return OpenaiImageResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errorMsg ErrorMsg
		if err := json.Unmarshal(respBody, &errorMsg); err != nil {
			return OpenaiImageResponse{}, err
		}
		return OpenaiImageResponse{}, errors.New(errorMsg.Error.Message)
	}

	if err := json.Unmarshal(respBody, &openaiResponse); err != nil {
		return OpenaiImageResponse{}, err
	}

	return openaiResponse, nil
}
