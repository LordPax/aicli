package sdk

import (
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkTextInstance SdkTextInterface

type Sdk struct {
	apiUrl string
	ApiKey string
	Model  string
}

// type SdkImage struct {
// 	Dimenssions string
// }

func InitSdkText(sdkType string, apiKey string, model string) error {
	l := lang.GetLocalize()

	switch sdkType {
	case "openai":
		sdkTextInstance = NewOpenaiText(apiKey, model)
		return nil
	default:
		return fmt.Errorf(l.Get("unknown-sdk"), sdkType)
	}
}

func InitSdk() error {
	textType := config.CONFIG_INI.Section("text").Key("type").String()
	textApiKey := config.CONFIG_INI.Section("text").Key("api_key").String()
	textModel := config.CONFIG_INI.Section("text").Key("model").String()
	if err := InitSdkText(textType, textApiKey, textModel); err != nil {
		return err
	}

	// TODO : init sdk for image, audio and translate

	return nil
}

func GetSdkText() SdkTextInterface {
	return sdkTextInstance
}
