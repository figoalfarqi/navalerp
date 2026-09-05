package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID      string `json:"user_id,omitempty"`
	AppUserID   int    `json:"app_user_id"`
	Username    string `json:"username"`
	AppRoleID   int    `json:"app_role_id"`
	Role        string `json:"role,omitempty"`
	FullName    string `json:"app_user_name,omitempty"`
	PhotoURL    string `json:"app_user_photo_url,omitempty"`
	StatusID    int    `json:"app_user_status_id,omitempty"`
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

func GenerateNavalToken(secret string, userID string, userName, fullName, role string, authVersion int, ttl time.Duration) (string, error) {
	claims := &Claims{
		UserID:      userID,
		AppUserID:   1,
		Username:    userName,
		AppRoleID:   5, // SuperAdmin / Admin
		Role:        role,
		FullName:    fullName,
		StatusID:    1,
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
