package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"modules/common/server/httpdata"
	"modules/common/server/httpdata/d7errors"
	"modules/common/server/httpdata/d7errors/codes"
	"time"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// JSON이 두번 쓰이는 것을 대비해서 Body 확인
		isBodyWritten := ctx.Writer.Written()
		err := ctx.Errors.Last()

		if err != nil {
			var netErr *d7errors.NetError
			// Get httpdata from NetError
			if errors.As(err, &netErr) {
				code := netErr.GetCode()
				res := httpdata.NewBaseResponse(code, nil)
				// Abort with the appropriate status code and domain
				if !isBodyWritten {
					res.Error(ctx)
				}
			} else {
				res := httpdata.NewBaseResponse(codes.GlobalInternalServerError, nil)
				if !isBodyWritten {
					res.Error(ctx)
				}
			}

		}
	}
}

func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
		c.Header("Last-Modified", time.Now().String())
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "-1")
		c.Next()
	}

}
