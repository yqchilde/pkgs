package ecode

import "github.com/yqchilde/gin-skeleton/pkg/errcode"

var (
	ErrEmailOrPassword   = errcode.NewError(20101, "邮箱或密码错误")
	ErrRegisterFailed    = errcode.NewError(20102, "注册失败")
	ErrDeleteApplication = errcode.NewError(20103, "应用删除失败")
)
