package d7errors

import (
	"github.com/sirupsen/logrus"
	"modules/common/server/httpdata/d7errors/codes"
)

type NetError struct {
	Code codes.WebCode `json:"code"`
	Err  error         `json:"error"`
}

func NewNetError(code codes.WebCode, err error) *NetError {
	logrus.Errorf("Code : %d, Error : %v", code, err)
	if err != nil {
		return &NetError{Code: code, Err: err}
	}
	return &NetError{Code: code, Err: nil}
}

func (e *NetError) GetCode() codes.WebCode {
	return e.Code
}
func (e *NetError) Error() string {
	return e.Err.Error()
}
