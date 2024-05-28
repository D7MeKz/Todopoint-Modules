package d7jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// GetBasic extracts the basic token from the Authorization.
// It returns base64 encoded token.
func GetBasic(ctx *gin.Context) (string, error) {
	auth, err := getAuthorization(ctx)
	if err != nil {
		return "", err
	}

	token, err := splitBasic(auth)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetBearerToken extracts the bearer token from the Authorization header
func GetBearerToken(ctx *gin.Context) (string, error) {
	auth, err := getAuthorization(ctx)
	if err != nil {
		return "", err
	}

	token, err := splitBearer(auth)
	if err != nil {
		return "", err
	}
	return token, nil
}

func getAuthorization(ctx *gin.Context) (string, error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		return "", errors.New("authorization value is Empty")
	}
	return auth, nil
}

func splitBasic(token string) (string, error) {
	if len(token) < 6 {
		return "", errors.New("invalid Authorization value")
	}
	return token[6:], nil
}

func splitBearer(token string) (string, error) {

	if len(token) < 7 {
		return "", errors.New("invalid Authorization value")
	}
	return token[7:], nil
}
