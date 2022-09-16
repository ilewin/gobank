package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	minSecretKeySize  = 32
	ErrMalformedToken = fmt.Errorf("malformed token")
	ErrExpiredToken   = fmt.Errorf("expired token")
)

type JWTMaker struct {
	key string
}

func NewJWTMaker(key string) (Maker, error) {
	if len(key) < minSecretKeySize {
		return nil, fmt.Errorf("The key is to short, minimum %d characters", minSecretKeySize)
	}
	return &JWTMaker{key: key}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(m.key))

}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.key), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		} else {
			return nil, ErrMalformedToken
		}
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("unexpected payload type")
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
