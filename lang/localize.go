package lang

var local *Localize

type LangString map[string]string

type Localize struct {
	lang    string
	strings map[string]*LangString
}

func NewLocalize() *Localize {
	return &Localize{
		lang:    "en_US.UTF-8",
		strings: make(map[string]*LangString),
	}
}

func GetLocalize() *Localize {
	if local == nil {
		local = NewLocalize()
	}
	return local
}

func (l *Localize) SetLang(lang string) {
	l.lang = lang
}

func (l *Localize) AddStrings(strings *LangString, lang ...string) {
	for _, la := range lang {
		l.strings[la] = strings
	}
}

func (l *Localize) Get(key string) string {
	lang, ok := l.strings[l.lang]

	if !ok {
		return (*l.strings["en_US.UTF-8"])[key]
	}

	return (*lang)[key]
}
