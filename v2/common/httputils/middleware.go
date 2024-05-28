package httputils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// JSON이 두번 쓰이는 것을 대비해서 Body 확인
		isBodyWritten := ctx.Writer.Written()
		err := ctx.Errors.Last()

		if err != nil {
			netErr := &NetError{}
			// Get httpdata from NetError
			if errors.As(err, &netErr) {
				res := NewErrorBaseResponse(netErr)
				// Abort with the appropriate status code and domain
				if !isBodyWritten {
					res.Failed(ctx)
				}
			} else {
				res := NewErrorBaseResponse(netErr)
				if !isBodyWritten {
					res.Failed(ctx)
				}
			}
		}
	}
}
