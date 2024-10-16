package sdk

import (
	"errors"
	"fmt"

	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
)

var sdkImageInstance IImageService

type IImageService interface {
	ISdk
	ISdkImage
}

type ISdkImage interface {
	SendRequest(prompt string) error
	SetModel(model string)
	GetModel() string
	SetSize(size string)
	GetSize() string
	SetImageNb(imageNb int)
	GetImageNb() int
	SetOutput(output string)
	GetOutput() string
}

type SdkImage struct {
	Model   string
	Size    string
	ImageNb int
	Output  string
}

func InitSdkImage(sdk string) error {
	var err error

	l := lang.GetLocalize()
	sdkType, apiKey, err := getConfigImage(sdk)
	if err != nil {
		return err
	}

	switch sdkType {
	case "openai":
		sdkImageInstance, err = NewOpenaiImage(apiKey)
	default:
		return fmt.Errorf(l.Get("unknown-sdk"), sdk)
	}

	if err != nil {
		return err
	}

	return nil
}

func getConfigImage(sdkType string) (string, string, error) {
	l := lang.GetLocalize()
	configImage := config.CONFIG_INI.Section("image")

	if sdkType == "" {
		sdkType = configImage.Key("type").String()
		if sdkType == "" {
			return "", "", errors.New(l.Get("type-required"))
		}
	}

	apiKey := configImage.Key("apiKey").String()
	if apiKey == "" {
		apiKey = configImage.Key(sdkType + "-apiKey").String()
	}

	return sdkType, apiKey, nil
}

func GetSdkImage() IImageService {
	return sdkImageInstance
}

func SetSdkImage(s IImageService) {
	sdkImageInstance = s
}

func (s *SdkImage) SetModel(model string) {
	s.Model = model
}

func (s *SdkImage) GetModel() string {
	return s.Model
}

func (s *SdkImage) SetSize(size string) {
	s.Size = size
}

func (s *SdkImage) GetSize() string {
	return s.Size
}

func (s *SdkImage) SetImageNb(imageNb int) {
	s.ImageNb = imageNb
}

func (s *SdkImage) GetImageNb() int {
	return s.ImageNb
}

func (s *SdkImage) SetOutput(output string) {
	s.Output = output
}

func (s *SdkImage) GetOutput() string {
	return s.Output
}
