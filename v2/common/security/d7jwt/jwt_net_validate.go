package d7jwt

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/netservice"
)

// Validate is a function to validate token
// They request to auth service to validate token
func Validate(ctx *gin.Context) (int, error) {
	token, err := getAuthorization(ctx)
	if err != nil {
		return -1, err
	}

	// Get request to auth service
	data, err := netservice.Get("http://localhost:3001/auth/valid", token)
	if err != nil {
		return -1, err
	}

	// Unmarshal data to UserId
	var resp httputils.BaseResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return -1, err
	}

	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		return -1, errors.New("invalid data type")
	}

	// Convert map to JSON for unmarshaling into AuthData
	dataBytes, err := json.Marshal(dataMap)
	if err != nil {
		return -1, err
	}

	var authData UserId
	err = json.Unmarshal(dataBytes, &authData)
	if err != nil {
		return -1, err
	}

	return authData.Id, nil
}
