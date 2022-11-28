package auth

import (
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func NewCasbin(db *gorm.DB, modelPath string) (*casbin.SyncedEnforcer, error) {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			fmt.Printf("[Auth] NewCasbin gormadapter.NewAdapterByDB failed: %v", err)
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(modelPath, adapter)
		syncedEnforcer.AddFunction("ParamsMatch", paramsMatchFunc)
	})
	err := syncedEnforcer.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("[Auth] NewCasbin syncedEnforcer.LoadPolicy failed: %v", err)
	}
	return syncedEnforcer, nil
}

func paramsMatch(key1, key2 string) bool {
	key1 = strings.Split(key1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

func paramsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return paramsMatch(name1, name2), nil
}
