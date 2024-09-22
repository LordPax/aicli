package sdk

type ISdk interface {
	GetName() string
}

type Sdk struct {
	Name   string
	ApiUrl string
	ApiKey string
}

func (s *Sdk) GetName() string {
	return s.Name
}

func InitSdk() error {
	if err := InitSdkText(""); err != nil {
		return err
	}

	if err := InitSdkTranslate(""); err != nil {
		return err
	}

	// TODO : init sdk for image, audio

	return nil
}
