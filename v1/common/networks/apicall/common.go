package apicall

import (
	"encoding/json"
	"fmt"
	"modules/common/server/httpdata/d7errors/codes"
	"net/http"
	"time"
)

type ReqClient struct {
	Client *http.Client
}

func NewReqClient() *ReqClient {
	s, _ := time.ParseDuration("60s")
	return &ReqClient{
		Client: &http.Client{
			Timeout: s,
		},
	}
}

func parseBody(response *http.Response) (*ExternalInfo, error) {
	statusCode := response.StatusCode

	// Convert to empty interface
	res := make(map[string]interface{})
	err := json.NewDecoder(response.Body).Decode(&res)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	_ = response.Body.Close()

	// Extract ExternalInfo from domain
	exInfo := extractFrom(res, statusCode)

	return exInfo, nil
}

type ExternalInfo struct {
	Code   codes.WebCode
	Status int
}

func (e *ExternalInfo) IsSuccess() bool {
	if e.Status == 200 || e.Status == 201 {
		return true
	}
	return false
}

func extractFrom(res map[string]interface{}, status int) *ExternalInfo {
	// Convert to WebCode
	code, ok := res["code"]
	if !ok {
		return nil
	}
	convertedCode := codes.ConvertFrom(code)
	return &ExternalInfo{
		Code:   convertedCode,
		Status: status,
	}
}
