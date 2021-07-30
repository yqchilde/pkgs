package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
)

func (s *Service) AddCasbinRule(ctx context.Context, req *request.CasbinInReceive) error {
	return s.dao.AddCasbinRule(ctx, req.AuthorityID, req.CasbinInfos)
}

func (s *Service) CheckJwtIsBlackList(ctx context.Context, jwtToken string) (bool, error) {
	has, _, err := s.dao.GetByJwtToken(ctx, jwtToken)
	if err != nil {
		return false, errors.Wrap(err, "[Service.auth] CheckJwtIsBlackList failed")
	}

	return has, err
}

func (s *Service) AddJwtBlacklist(ctx context.Context, jwtToken string) error {
	err := s.dao.AddJwtBlacklist(ctx, &model.JwtBlacklist{Jwt: jwtToken})
	if err != nil {
		return errors.Wrap(err, "[Service.auth] AddJwtBlacklist failed")
	}

	return nil
}
