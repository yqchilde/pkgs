package application

import "github.com/yqchilde/gin-skeleton/pkg/app"

var response = app.NewResponse()

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	AppName string `json:"app_name" binding:"required"`
}
