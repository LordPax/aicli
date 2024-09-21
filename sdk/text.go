package sdk

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
	// Source struct {
	// 	Type      string `json:"type"`
	// 	MediaType string `json:"media_type"`
	// 	Data      string `json:"data"`
	// } `json:"source"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

func (m *Message) GetContent() string {
	var text string

	if len(m.Content) == 1 {
		return m.Content[0].Text
	}

	for _, c := range m.Content {
		if c.Type != "text" {
			continue
		}
		text += "\n" + c.Text + "\n"
	}

	return text
}

type ErrorMsg struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type ITextService interface {
	ISdk
	ISdkText
	ITextHistory
}

type ISdkText interface {
	SetTemp(temp float64)
	GetTemp() float64
}

type SdkText struct {
	Temp float64
}

func (s *SdkText) SetTemp(temp float64) {
	s.Temp = temp
}

func (s *SdkText) GetTemp() float64 {
	return s.Temp
}
