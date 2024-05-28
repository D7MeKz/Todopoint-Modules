package netservice

import (
	"errors"
	"io"
	"net/http"
)

// Get is the http request to other service.
// If you want to unmarshall the response, you could use json.Unmarshal
func Get(url string, token string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set header
	req.Header.Add("Authorization", token)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Check request success
	ok := isSuccess(resp)
	if !ok {
		return nil, errors.New("request to external request is failed")
	}
	// read body to bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func isSuccess(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
