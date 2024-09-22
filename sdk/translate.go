package sdk

import (
	"errors"
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkTranslateInstance ITranslateService

type ITranslateService interface {
	ISdk
	ISdkTranslate
}

type ISdkTranslate interface {
	SendRequest(text string) (string, error)
	SetSourceLang(source string)
	GetSourceLang() string
	SetTargetLang(target string)
	GetTargetLang() string
}

type SdkTranslate struct {
	SourceLang string
	TargetLang string
}

func InitSdkTranslate(sdk string) error {
	var err error

	l := lang.GetLocalize()
	sdkType, apiKey, err := getConfigTranslate(sdk)
	if err != nil {
		return err
	}

	switch sdkType {
	case "deepl":
		sdkTranslateInstance, err = NewDeepL(apiKey)
	default:
		return fmt.Errorf(l.Get("unknown-sdk"), sdkType)
	}

	if err != nil {
		return err
	}

	return nil
}

func getConfigTranslate(sdkType string) (string, string, error) {
	l := lang.GetLocalize()
	configTranslate := config.CONFIG_INI.Section("translate")

	if sdkType == "" {
		sdkType = configTranslate.Key("type").String()
		if sdkType == "" {
			return "", "", errors.New(l.Get("type-required"))
		}
	}

	apiKey := configTranslate.Key("apiKey").String()
	if apiKey == "" {
		apiKey = configTranslate.Key(sdkType + "-apiKey").String()
		if apiKey == "" {
			return "", "", fmt.Errorf(l.Get("api-key-required"), sdkType)
		}
	}

	return sdkType, apiKey, nil
}

func GetSdkTranslate() ITranslateService {
	return sdkTranslateInstance
}

func SetSdkTranslate(s ITranslateService) {
	sdkTranslateInstance = s
}

func (s *SdkTranslate) SetSourceLang(source string) {
	s.SourceLang = source
}

func (s *SdkTranslate) GetSourceLang() string {
	return s.SourceLang
}

func (s *SdkTranslate) SetTargetLang(target string) {
	s.TargetLang = target
}

func (s *SdkTranslate) GetTargetLang() string {
	return s.TargetLang
}
