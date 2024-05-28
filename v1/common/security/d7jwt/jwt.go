package d7jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	secret = "todopoint123"
)

// TokenClaims is a struct for JWT token claims.
type TokenClaims struct {
	// TokenUUID is a unique identifier for the token.
	TokenUUID string `json:"tid"`
	// UserID is the user ID of the user.
	UserID int `json:"user_id"`
	// Email is the email of the user.
	Email string `json:"email"`
	// Role is the role of the user.
	Role []string `json:"role"`
	// MapClaims is the map of the claims. It contains iss, exp.
	jwt.MapClaims
}

func NewTokenClaims(uid int) *TokenClaims {
	claim := TokenClaims{
		TokenUUID: uuid.NewString(),
		UserID:    uid,
		MapClaims: jwt.MapClaims{
			"iss": "d7mekz",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return &claim
}

// Generate generates a token from the claims.
func (t TokenClaims) Generate() (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	token, err := claim.SignedString([]byte(secret))
	return token, err
}

// IsExpired checks the token is expired or not.
func IsExpired(token string) (bool, error) {
	_, err := validate(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return true, nil
		} else {
			return false, err
		}
	}

	return false, nil
}

// validate validates the token that is expired or not.
// If the token is valid, it returns the claims.
func validate(token string) (*TokenClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	// Validate Token Claim type.
	if claims, ok := parsed.Claims.(*TokenClaims); ok && parsed.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// GetIdFrom extracts the user ID from the token.
func GetIdFrom(token string) (int, error) {
	claim, err := validate(token)
	if err != nil {
		return -1, err
	}
	return claim.UserID, nil
}

func GetIdFromHeader(ctx *gin.Context) (int, error) {
	token, err := GetBearerToken(ctx)
	if err != nil {
		return -1, err
	}
	return GetIdFrom(token)
}
