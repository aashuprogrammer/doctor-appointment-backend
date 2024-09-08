package token

import "time"

type Maker interface {
	CreateAccessToken(id int64, name string, email string, role string, duration time.Duration) (string, *AccessTokenPayload, error)
	VerifyAccessToken(token string) (*AccessTokenPayload, error)
}
