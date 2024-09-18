package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
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
	if defaultVal != "" {
		prompt = fmt.Sprintf("[%s] %s", defaultVal, prompt)
	}

	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	if text == "" && defaultVal != "" {
		return defaultVal
	}

	if text == "" && !nullable {
		return Input(prompt, defaultVal, nullable)
	}

	return text
}
