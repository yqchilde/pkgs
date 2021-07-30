package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
	"github.com/yqchilde/gint/internal/model/response"
	"github.com/yqchilde/gint/pkg/app"
	"github.com/yqchilde/gint/pkg/auth"
)

var (
	ErrUserExists    = errors.New("the user exists")
	ErrUserNotExists = errors.New("the user not exists")
)

func (s *Service) SignUp(ctx context.Context, req *request.SignUp) error {
	// Check user exists
	has, _, err := s.dao.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.Wrap(err, "[Service.public] SignUp dao.GetByEmail failed")
	}
	if has {
		return ErrUserExists
	}

	// Encryption password
	pwd, err := auth.HashAndSalt(req.Password)
	if err != nil {
		return errors.Wrap(err, "[Service.public] SignUp auth.HashAndSalt failed")
	}

	u := &model.User{
		UserID:      uuid,
		Email:       req.Email,
		Username:    req.Username,
		Password:    pwd,
		AuthorityID: req.AuthorityID,
	}
	if err := s.dao.CreateUser(ctx, u); err != nil {
		return errors.Wrap(err, "[Service.public] dao.CreateUser failed")
	}

	return nil
}

func (s *Service) SignIn(ctx context.Context, req *request.SignIn) (*response.SignIn, error) {
	// Check user exists
	has, user, err := s.dao.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.Wrap(err, "[Service.public] SignIn dao.GetByEmail failed")
	}
	if !has {
		return nil, ErrUserNotExists
	}

	// Verify password
	if !auth.ComparePasswords(user.Password, req.Password) {
		return nil, errors.Wrap(err, "[Service.public] SignIn auth.ComparePasswords failed")
	}

	// Sign jwt
	payload := map[string]interface{}{
		"user_id":      user.UserID,
		"username":     user.Username,
		"authority_id": user.AuthorityID,
	}
	tokenStr, err := app.Sign(ctx, payload, s.cfg.Server.JwtSecret, s.cfg.Server.JwtExpireTime)
	if err != nil {
		return nil, errors.Wrap(err, "[Service.public] SignIn app.Sign failed")
	}

	resp := &response.SignIn{
		User:      user,
		Token:     tokenStr,
		ExpiresAt: s.cfg.Server.JwtExpireTime,
	}

	return resp, nil
}
