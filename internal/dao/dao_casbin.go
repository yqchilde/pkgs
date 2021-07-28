package dao

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
	"github.com/yqchilde/gint/pkg/auth"
	"github.com/yqchilde/gint/pkg/logger"
)

func (d *Dao) newCasbin() *casbin.SyncedEnforcer {
	return auth.NewCasbin(d.db, d.cfg.Server.CasbinModelPath)
}

func (d *Dao) clearCasbin(v int, p ...string) bool {
	e := d.newCasbin()
	status, err := e.RemoveFilteredPolicy(v, p...)
	if err != nil {
		logger.Error("[Dao.casbin] clearCasbin failed: %s", err.Error())
	}
	return status
}

func (d *Dao) UpdateCasbin(ctx context.Context, authorityID string, casbinInfos []request.CasbinInfo) error {
	// remove casbin policy
	d.clearCasbin(0, authorityID)

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
	e := d.GetCasbin()
	status, err := e.AddPolicies(rules)
	if err != nil {
		return errors.Wrap(err, "[repo.auth] CasbinUpdate AddPolicies failed")
	}
	if !status {
		return errors.New("存在相同API，添加失败，请联系管理员")
	}

	return nil
}