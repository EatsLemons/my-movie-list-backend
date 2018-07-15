package auth

import (
	"time"
)

type JWT struct {
	secret        string
	tokenDuration time.Duration
}

func NewJWT(secret string, tokenDuration time.Duration) *JWT {
	res := JWT{
		secret:        secret,
		tokenDuration: tokenDuration,
	}
	return &res
}
