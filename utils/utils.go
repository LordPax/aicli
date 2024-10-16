package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/h2non/filetype"
	"golang.org/x/term"
)

const (
	Escape = "\x1b"
	Reset  = Escape + "[0m"
	Red    = Escape + "[31m"
	Green  = Escape + "[32m"
	Blue   = Escape + "[34m"
)

var (
	IMAGE = []string{"jpg", "jpeg", "png", "gif", "webp"}
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

func IsFileType(buf []byte, fileType []string) string {
	for _, t := range fileType {
		if filetype.Is(buf, t) {
			return t
		}
	}
	return ""
}

func GetFileFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func ConvertB64ToImage(name, b64 string) error {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}

	return os.WriteFile(name, dec, 0644)
}
