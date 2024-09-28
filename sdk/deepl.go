package sdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LordPax/aicli/utils"
)

type DeepLBody struct {
	Text       []string `json:"text"`
	TargetLang string   `json:"target_lang"`
	SourceLang string   `json:"source_lang"`
}

type DeepLResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

type DeepLError struct {
	Message string `json:"message"`
}

type DeepL struct {
	Sdk
	SdkTranslate
}

func NewDeepL(apiKey string) (*DeepL, error) {
	return &DeepL{
		Sdk: Sdk{
			Name:   "deepl",
			ApiUrl: "https://api-free.deepl.com/v2/translate",
			ApiKey: apiKey,
			Inerte: false,
		},
		SdkTranslate: SdkTranslate{
			SourceLang: "",
			TargetLang: "",
		},
	}, nil
}

func (d *DeepL) SendRequest(text string) (string, error) {
	var deeplResponse DeepLResponse

	jsonBody, err := json.Marshal(DeepLBody{
		Text:       []string{text},
		TargetLang: d.GetTargetLang(),
		SourceLang: d.GetSourceLang(),
	})
	if err != nil {
		return "", err
	}

	resp, err := utils.PostRequest(d.ApiUrl, jsonBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "DeepL-Auth-Key " + d.ApiKey,
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		var deepLError DeepLError
		if err := json.Unmarshal(respBody, &deepLError); err != nil {
			return "", err
		}
		return "", errors.New(deepLError.Message)
	}

	if err := json.Unmarshal(respBody, &deeplResponse); err != nil {
		return "", err
	}

	return deeplResponse.Translations[0].Text, nil
}
