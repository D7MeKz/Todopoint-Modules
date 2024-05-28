package httputils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"modules/v2/common/httputils/codes"
)

// BaseResponse is common response that use in success and failed
type BaseResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (b BaseResponse) String() string {
	data, _ := json.Marshal(b.Data)
	return fmt.Sprintf("{Status:%s Data:%s}", b.Status, string(data))
}

func NewSuccessBaseResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		Status: "Success",
		Data:   data,
	}
}

// OKSuccess uses when status code is 200
func (b BaseResponse) OKSuccess(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(200, b)
}

// CreatedSuccess uses when status code is 201
func (b BaseResponse) CreatedSuccess(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(201, b)
}

// ErrorResponse is the error response format
type ErrorResponse struct {
	Code    codes.ErrorCode `json:"code"`
	Message string          `json:"message"`
}

func NewErrorBaseResponse(data *NetError) *BaseResponse {
	logrus.Errorf("Error: %v", data)
	res := ErrorResponse{Code: data.Code, Message: codes.GetErrorMsg(data.Code)}
	return &BaseResponse{
		Status: "Error",
		Data:   res,
	}
}

// Failed response failed formatted response.
// It converts NetError to ErrorResponse to extract necessary things.
func (b BaseResponse) Failed(ctx *gin.Context) {
	res := b.GetErrorData()
	status := codes.ParseStatusCode(res.Code)
	ctx.AbortWithStatusJSON(status, b)
	return
}

func (b BaseResponse) GetErrorData() *ErrorResponse {
	if c, ok := b.Data.(ErrorResponse); ok {
		return &c
	}
	return nil
}

// UnmarshalErrorData unmarshal error data to ErrorResponse
func (b BaseResponse) UnmarshalErrorData() *ErrorResponse {
	data, _ := json.Marshal(b.Data)
	var res ErrorResponse
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil
	}
	return &res
}
