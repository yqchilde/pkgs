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
// @Summary developer create application
// @Description developer create application in application center
// @Tags App
// @Produce json
// @Param req body CreateRequest true "Request parameter"
// @Success 200 {object} app.Response
// @Router /app/v1/application [post]
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
		if _, ok := err.(*errcode.Error); ok {
			response.Error(ctx, err)
			return
		}
		response.Error(ctx, ecode.ErrEmailOrPassword)
		return
	}

	response.Success(ctx, data)
}

// DeleteApp developer delete application
// @Summary developer delete application
// @Description developer delete application in application center
// @Tags App
// @Produce json
// @Param id query string true "application id"
// @Success 200 {object} app.Response
// @Router /app/v1/application/:id [delete]
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
		if _, ok := err.(*errcode.Error); ok {
			response.Error(ctx, err)
			return
		}
		response.Error(ctx, ecode.ErrDeleteApplication)
		return
	}

	response.Success(ctx, nil)
}
