package internal

import (
	"net/http"
	"os"
)

func HttpRequest(url string, bodyFilePath string, httpMethod string, apiKey string) (*http.Response, error) {
	file, err := os.Open(bodyFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	req, err := http.NewRequest(httpMethod, url, file)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"api-key":      []string{apiKey},
	}

	return http.DefaultClient.Do(req)
}
