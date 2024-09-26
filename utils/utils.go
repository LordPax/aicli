package utils

import (
	"bytes"
	"net/http"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
	Escape = "\x1b"
	Reset  = Escape + "[0m"
	Red    = Escape + "[31m"
	Green  = Escape + "[32m"
	Blue   = Escape + "[34m"
)

func FileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func RmTmpDir() error {
	return os.RemoveAll(os.Getenv("TMP_DIR"))
}

func PostRequest(url string, data []byte, option map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range option {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Input(prompt string, defaultVal string, nullable bool) string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	t := term.NewTerminal(os.Stdin, prompt)

	line, _ := t.ReadLine()
	line = strings.TrimSpace(line)

	if line == "" && defaultVal != "" {
		return defaultVal
	}

	if line == "" && !nullable {
		return Input(prompt, defaultVal, nullable)
	}

	return line
}
