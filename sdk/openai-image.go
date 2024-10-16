package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/LordPax/aicli/lang"
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
	Data []struct {
		Url     string `json:"url"`
		B64Json string `json:"b64_json"`
	} `json:"data"`
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

func (o *OpenaiImage) SendRequest(prompt string) error {
	var openaiResponse OpenaiImageResponse
	format := "url"

	if o.GetOutput() != "" {
		format = "b64_json"
	}

	jsonBody, err := json.Marshal(OpenaiImageBody{
		Model:          o.GetModel(),
		Prompt:         prompt,
		N:              o.GetImageNb(),
		Size:           o.GetSize(),
		ResponseFormat: format,
	})
	if err != nil {
		return err
	}

	resp, err := utils.PostRequest(o.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.ApiKey,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errorMsg ErrorMsg
		if err := json.Unmarshal(respBody, &errorMsg); err != nil {
			return err
		}
		return errors.New(errorMsg.Error.Message)
	}

	if err := json.Unmarshal(respBody, &openaiResponse); err != nil {
		return err
	}

	if o.GetOutput() != "" {
		if err := convertFile(openaiResponse, o.GetOutput()); err != nil {
			return err
		}
		return nil
	}

	for _, image := range openaiResponse.Data {
		println(image.Url)
	}

	return nil
}

func convertFile(openaiResponse OpenaiImageResponse, output string) error {
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if (len(openaiResponse.Data)) == 1 {
		if err := utils.ConvertB64ToImage(output, openaiResponse.Data[0].B64Json); err != nil {
			return err
		}
		log.Printf(l.Get("image-save"), output)
		return nil
	}

	for i, image := range openaiResponse.Data {
		name := fmt.Sprintf("%s-%d", output, i)
		if err := utils.ConvertB64ToImage(name, image.B64Json); err != nil {
			return err
		}
		log.Printf(l.Get("image-save"), name)
	}

	return nil
}
