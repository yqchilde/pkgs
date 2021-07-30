package service

import (
	"github.com/yqchilde/gint/internal/dao"
	"github.com/yqchilde/gint/pkg/config"
	"github.com/yqchilde/gint/pkg/generator"
)

var PubSvc *Service
var AuthSvc *Service

var uuid = generator.NewUUID()

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
	PubSvc = s
	return s
}
