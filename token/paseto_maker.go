package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	paseto "github.com/o1egl/paseto"
)

var (
	ErrInvalidKey = fmt.Errorf("invalid key")
)

type PasetoMaker struct {
	passeto     *paseto.V2
	symetricKey []byte
}

func NewPasetoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, ErrInvalidKey
	}

	return &PasetoMaker{
		passeto:     paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}, nil

}

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return m.passeto.Encrypt(m.symetricKey, payload, nil)

}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload *Payload
	err := m.passeto.Decrypt(token, m.symetricKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
