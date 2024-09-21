package sdk

import (
	"errors"
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkTextInstance ITextService

type ISdk interface {
	SendRequest(text string) (Message, error)
	SetModel(model string)
	GetModel() string
	GetName() string
}

type Sdk struct {
	Name   string
	ApiUrl string
	ApiKey string
	Model  string
}

func (s *Sdk) GetName() string {
	return s.Name
}

func (s *Sdk) SetModel(model string) {
	s.Model = model
}

func (s *Sdk) GetModel() string {
	return s.Model
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
		if apiKey == "" {
			return "", "", "", 0, errors.New(l.Get("api-key-required"))
		}
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

func InitSdk() error {
	if err := InitSdkText(""); err != nil {
		return err
	}

	// TODO : init sdk for image, audio and translate

	return nil
}

func GetSdkText() ITextService {
	return sdkTextInstance
}

func SetSdkText(s ITextService) {
	sdkTextInstance = s
}
