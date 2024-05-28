package httpdata

import (
	"github.com/gin-gonic/gin"
	"modules/common/server/httpdata/d7errors/codes"
)

type SuccessResponse struct {
	Code    codes.WebCode `json:"code"`
	Message string        `json:"message"`
}

func NewSuccessResponse(code codes.WebCode) *SuccessResponse {
	return &SuccessResponse{Code: code, Message: "Success"}
}

type SuccessResponseWith struct {
	Code    codes.WebCode `json:"code"`
	Message string        `json:"message"`
	Data    any           `json:"httpdata"`
}

func NewSuccessResponseWith(code codes.WebCode, data any) *SuccessResponseWith {
	return &SuccessResponseWith{Code: code, Message: "Success", Data: data}
}

func SuccessWith(ctx *gin.Context, code codes.WebCode, data any) {
	status := codes.GetStatus(code)
	res := NewSuccessResponseWith(code, data)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, res)
}
func Success(ctx *gin.Context, code codes.WebCode) {
	status := codes.GetStatus(code)
	ctx.Header("Content-Type", "application/json")

	res := NewSuccessResponse(code)
	ctx.JSON(status, res)
}

type BaseResponse struct {
	Code codes.WebCode `json:"code"`
	Data interface{}   `json:"data"`
}

func NewBaseResponse(code codes.WebCode, data interface{}) *BaseResponse {
	return &BaseResponse{Code: code, Data: data}
}

func (b BaseResponse) Success(ctx *gin.Context) {
	status := codes.GetStatus(b.Code)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, b)
}

func (b BaseResponse) Error(ctx *gin.Context) {
	status := codes.GetStatus(b.Code)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, b)
}
