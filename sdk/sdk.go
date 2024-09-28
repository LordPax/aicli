package sdk

type ISdk interface {
	GetName() string
	SetInerte(bool)
	GetInerte() bool
}

type Sdk struct {
	Name   string
	ApiUrl string
	ApiKey string
	Inerte bool
}

func (s *Sdk) SetInerte(inerte bool) {
	s.Inerte = inerte
}

func (s *Sdk) GetInerte() bool {
	return s.Inerte
}

func (s *Sdk) GetName() string {
	return s.Name
}
