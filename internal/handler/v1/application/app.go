package application

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/internal/ecode"
	"github.com/yqchilde/gin-skeleton/internal/service"
	"github.com/yqchilde/gin-skeleton/pkg/app"
	"github.com/yqchilde/gin-skeleton/pkg/errcode"
	"github.com/yqchilde/gin-skeleton/pkg/log"
)

// CreateApp developer create application
func CreateApp(ctx *gin.Context) {
	var req CreateRequest
	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		log.Warnf("CreateApp bind and validate param err: %v", errs)
		response.Error(ctx, errcode.ErrInvalidParam.WithDetails(errs.Errors()...))
		return
	}

	data, err := service.NewApplicationService(ctx).CreateApp(req.UserID, req.AppName)
	if err != nil {
		log.Warnf("CreateApp handler err: %v", err)
		response.Error(ctx, ecode.ErrEmailOrPassword)
		return
	}

	response.Success(ctx, data)
}

func DeleteApp(ctx *gin.Context) {
	appID := ctx.Param("id")
	if appID == "" {
		log.Warnf("DeleteApp param is empty: %v", appID)
		response.Error(ctx, errcode.ErrInvalidParam)
		return
	}

	err := service.NewApplicationService(ctx).DeleteApp(appID)
	if err != nil {
		log.Warnf("DeleteApp handler err: %v", err)
		response.Error(ctx, ecode.ErrDeleteApplication)
		return
	}

	response.Success(ctx, nil)
}
