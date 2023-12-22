package user

import (
	"context"
	"fmt"

	"github.com/seivanov1986/gocart/helpers"
)

type AuthLoginIn struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

func (s *service) Login(ctx context.Context, in AuthLoginIn) (*string, error) {
	password, err := s.hub.User().PasswordByLogin(ctx, in.Login)
	if err != nil {
		return nil, err
	}

	if !helpers.PasswordHashVerify(in.Password, *password) {
		return nil, fmt.Errorf("user has not found")
	}

	sessionId := helpers.GenerateHashSession()
	err = s.sessionManager.Set(sessionId, 1, 0)
	if err != nil {
		return nil, err
	}

	return &sessionId, nil
}
