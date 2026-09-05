package sse

import "errors"

type Claims struct {
	UserID int
	Role   string
}

func ValidateJWT(token string) (*Claims, error) {
	// CONTOH SAJA (di real pakai jwt-go)
	if token == "driver-token" {
		return &Claims{UserID: 1, Role: "driver"}, nil
	}
	if token == "admin-token" {
		return &Claims{UserID: 99, Role: "admin"}, nil
	}
	return nil, errors.New("invalid token")
}
