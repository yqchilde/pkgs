package service

import (
	"github.com/yqchilde/gint/internal/dao"
	"github.com/yqchilde/gint/pkg/config"
)

var AuthSvc *Service

type Service struct {
	cfg *config.Config
	dao *dao.Dao
}

func New(cfg *config.Config, dao *dao.Dao) *Service {
	s := &Service{
		cfg: cfg,
		dao: dao,
	}
	AuthSvc = s
	return s
}
