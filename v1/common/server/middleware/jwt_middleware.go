package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"modules/common/security/d7jwt"
	"modules/common/server/httpdata/d7errors/codes"
	"modules/common/server/httpdata/d7success"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := d7jwt.GetBearerToken(ctx)
		if err != nil {
			logrus.Errorf("Invalid Authorization value: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Validate Token
		ok, err := d7jwt.IsExpired(token)
		if err != nil {
			logrus.Errorf("Invalid Token : %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// if token expired
		if ok {
			logrus.Error("Token expired")
			res := d7success.NewErrorResponse(codes.TokenExpired)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

	}
}
