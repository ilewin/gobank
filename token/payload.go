package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssuesAt time.Time `json:"issued_at"`
	Expired  time.Time `json:"expired_at"`
}

func (p *Payload) Valid() error {
	if p.Expired.Before(time.Now()) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:       id,
		Username: username,
		IssuesAt: time.Now(),
		Expired:  time.Now().Add(duration),
	}, nil
}
