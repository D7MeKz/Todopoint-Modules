package d7jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := GetBearerToken(ctx)
		if err != nil {
			logrus.Errorf("Invalid Authorization value: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Validate Token
		ok, err := IsExpired(token)
		if err != nil {
			logrus.Errorf("Invalid Token : %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// if token expired
		if ok {
			logrus.Error("Token expired")
			res := httputils.NewErrorBaseResponse(
				httputils.NewNetError(codes.TokenExpired, errors.New("token expired")))

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

	}
}
