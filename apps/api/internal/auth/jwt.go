package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	AppUserID   int    `json:"app_user_id"`
	Username    string `json:"username"`
	AppRoleID   int    `json:"app_role_id"`
	AuthVersion int    `json:"auth_version"`
	jwt.RegisteredClaims
}

func GenerateToken(secret string, appUserID int, userName string, appRoleID, authVersion int, ttl time.Duration) (string, error) {
	claims := &Claims{
		AppUserID:   appUserID,
		Username:    userName,
		AppRoleID:   appRoleID,
		AuthVersion: authVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString([]byte(secret))
}

func ParseToken(secret, tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
