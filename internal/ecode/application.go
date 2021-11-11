package ecode

import "github.com/yqchilde/gin-skeleton/pkg/errcode"

var (
	ErrEmailOrPassword   = errcode.NewError(20001, "application.email_or_password")
	ErrRegisterFailed    = errcode.NewError(20002, "application.register_failed")
	ErrDeleteApplication = errcode.NewError(20003, "application.delete_application")
)
