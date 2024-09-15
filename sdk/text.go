package sdk

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TextBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type TextResponse struct {
	Choices []struct {
		Index   int64   `json:"index"`
		Message Message `json:"message"`
	}
}

type SdkTextInterface interface {
	SendRequest(text string) (string, error)
	AppendHistory(text Message)
	SaveHistory() error
	LoadHistory() error
	GetHistory() []Message
}

type SdkText struct {
	History []Message
	Temp    float64
}

func (s *SdkText) AppendHistory(text Message) {
	s.History = append(s.History, text)
}

func (s *SdkText) SaveHistory() error {
	return nil
}

func (s *SdkText) LoadHistory() error {
	return nil
}

func (s *SdkText) GetHistory() []Message {
	return s.History
}
