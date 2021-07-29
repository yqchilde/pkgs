package dao

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/pkg/errors"

	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
	"github.com/yqchilde/gint/pkg/auth"
	"github.com/yqchilde/gint/pkg/logger"
)

func (d *Dao) newCasbin() *casbin.SyncedEnforcer {
	return auth.NewCasbin(d.db, d.cfg.Server.CasbinModelPath)
}

func (d *Dao) RemoveCasbinRule(v int, p ...string) bool {
	e := d.newCasbin()
	status, err := e.RemoveFilteredPolicy(v, p...)
	if err != nil {
		logger.Error("[Dao.casbin] RemoveCasbinRule failed: %s", err.Error())
	}
	return status
}

func (d *Dao) AddCasbinRule(ctx context.Context, authorityID string, casbinInfos []request.CasbinInfo) error {
	// remove casbin policy
	d.RemoveCasbinRule(0, authorityID)

	var rules [][]string
	for _, v := range casbinInfos {
		cm := model.Casbin{
			PType:       "p",
			AuthorityId: authorityID,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := d.newCasbin()
	_, err := e.AddPolicies(rules)
	if err != nil {
		return errors.Wrap(err, "[Dao.casbin] AddCasbinRule AddPolicies failed")
	}
	return nil
}
