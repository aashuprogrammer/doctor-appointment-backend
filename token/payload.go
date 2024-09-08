package token

import "time"

type AccessTokenPayload struct {
	Id       int64
	Name     string
	Email    string
	Duration time.Duration
	IssueAt  time.Time
	ExpireAt time.Time
}

func NewAccessTokenPayload(Id int64, Name string, Email string, Duration time.Duration) (*AccessTokenPayload, error) {
	return &AccessTokenPayload{Id: Id, Name: Name, Email: Email, Duration: Duration}, nil
}
