package httputils

import "modules/v2/common/httputils/codes"

// NetError have options and ErrorCode that contains status code.
type NetError struct {
	// options is metadata about service
	options netErrorOptions
	// Type is a unique data that contains http status code.
	Code codes.ErrorCode
	// Description is an error details.
	Description string
	// Err is an error message.
	Err error
}

func NewNetError(code codes.ErrorCode, err error) *NetError {
	return &NetError{
		options: opts,
		Code:    code,
		Err:     err,
	}
}

func (e *NetError) Error() string {
	return e.Err.Error()
}

func (e *NetError) Tag() string {
	return e.options.tag
}
