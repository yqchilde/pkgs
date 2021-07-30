package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
	"github.com/yqchilde/gint/internal/service"
	"github.com/yqchilde/gint/pkg/errcode"
	"github.com/yqchilde/gint/pkg/validator"
)

func SignUp(c *gin.Context) {
	var r request.SignUp
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Error(c, errcode.ErrBind)
		return
	}

	if err := validator.Verify(r, model.SignUpVerify); err != nil {
		response.Error(c, errcode.ErrParamValidation.WithDetails(err.Error()))
		return
	}

	err := service.PubSvc.SignUp(c, &r)
	if err != nil {
		response.Error(c, errcode.ErrInternalServerError.WithDetails(err.Error()))
		return
	}

	response.Success(c, nil)
}

func SignIn(c *gin.Context) {
	var r request.SignIn
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Error(c, errcode.ErrBind)
		return
	}

	if err := validator.Verify(r, model.SignInVerify); err != nil {
		response.Error(c, errcode.ErrParamValidation.WithDetails(err.Error()))
		return
	}

	data, err := service.PubSvc.SignIn(c, &r)
	if err != nil {
		response.Error(c, errcode.ErrInternalServerError.WithDetails(err.Error()))
		return
	}

	response.Success(c, data)
}
