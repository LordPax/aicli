package utils

import (
	"bytes"
	"net/http"
	"os"
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
