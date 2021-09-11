package application

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/internal/ecode"
	"github.com/yqchilde/gin-skeleton/internal/service"
	"github.com/yqchilde/gin-skeleton/pkg/app"
	"github.com/yqchilde/gin-skeleton/pkg/errcode"
	"github.com/yqchilde/gin-skeleton/pkg/log"
)

// Register developer register
// @Summary developer register
// @Description Only for developer register
// @Tags App
// @Produce json
// @Param req body RegisterRequest true "Request parameter"
// @Success 200 {object} app.Response
// @Router /app/v1/request [post]
func Register(ctx *gin.Context) {
	var req RegisterRequest
	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		log.Warnf("Register bind and validate param err: %v", errs)
		response.Error(ctx, errcode.ErrInvalidParam.WithDetails(errs.Errors()...))
		return
	}

	err := service.NewApplicationService(ctx).Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		log.Warnf("Register handler err: %v", err)
		response.Error(ctx, ecode.ErrRegisterFailed)
		return
	}

	response.Success(ctx, nil)
}

// Login developer login
// @Summary developer login
// @Description Only for developer login
// @Tags App
// @Produce json
// @Param req body LoginRequest true "Request parameter"
// @Success 200 {object} app.Response
// @Router /app/v1/login [post]
func Login(ctx *gin.Context) {
	var req LoginRequest
	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		log.Warnf("Login bind and validate param err: %v", errs)
		response.Error(ctx, errcode.ErrInvalidParam.WithDetails(errs.Errors()...))
		return
	}

	t, err := service.NewApplicationService(ctx).Login(req.Email, req.Password)
	if err != nil {
		log.Warnf("Login handler err: %v", err)
		response.Error(ctx, ecode.ErrEmailOrPassword)
		return
	}

	response.Success(ctx, t)
}
