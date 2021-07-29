package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/model/request"
	"github.com/yqchilde/gint/pkg/errcode"
	"github.com/yqchilde/gint/pkg/validator"
)

func AddCasbinRule(c *gin.Context) {
	var r request.CasbinInReceive
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Error(c, errcode.ErrBind)
		return
	}

	if err := validator.Verify(r, model.CasbinVerify); err != nil {
		response.Error(c, errcode.ErrParamValidation, err)
		return
	}

	return

	//err := service.AuthSvc.UpdateCasbin(c, &r)
	//if err != nil {
	//	logger.Warnf("Update casbin err: %v", err)
	//	response.ErrorWithMsg(c, err.Error())
	//	return
	//}
}
