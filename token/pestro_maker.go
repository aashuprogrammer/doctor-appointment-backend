package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

const (
	payloadKey = "payload"
)

type PestroMaker struct {
	symetricKey paseto.V4SymmetricKey
}

func NewPastroMaker(secretKey string) (Maker, error) {
	KeyByte := []byte(secretKey)
	V4symmetric, err := paseto.V4SymmetricKeyFromBytes(KeyByte)
	if err != nil {
		return nil, err
	}
	maker := &PestroMaker{
		symetricKey: V4symmetric,
	}
	return maker, err
}

func (maker *PestroMaker) CreateAccessToken(id int64, name string, email string, role string, duration time.Duration) (string, *AccessTokenPayload, error) {
	payload, err := NewAccessTokenPayload(id, name, email, duration)
	if err != nil {
		return "", nil, err
	}
	token := paseto.NewToken()
	token.SetExpiration(payload.ExpireAt)
	token.SetIssuedAt(payload.IssueAt)
	token.Set(payloadKey, payload)
	return token.V4Encrypt(maker.symetricKey, nil), payload, nil
}

func (maker *PestroMaker) VerifyAccessToken(tokenString string) (*AccessTokenPayload, error) {
	parse := paseto.NewParser()
	token, err := parse.ParseV4Local(maker.symetricKey, tokenString, nil)
	if err != nil {
		return nil, err
	}
	payload := AccessTokenPayload{}
	token.Get(payloadKey, &payload)
	payload.ExpireAt, _ = token.GetExpiration()
	payload.IssueAt, _ = token.GetIssuedAt()
	return &payload, nil
}
