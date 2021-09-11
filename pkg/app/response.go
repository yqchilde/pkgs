package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/pkg/errcode"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Details []string    `json:"details"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Msg:     errcode.Success.Msg(),
		Data:    data,
		Details: []string{},
	})
}

func (r *Response) Error(ctx *gin.Context, err error) {
	if err != nil {
		if v, ok := err.(*errcode.Error); ok {
			response := Response{
				Code:    v.Code(),
				Msg:     v.Msg(),
				Data:    gin.H{},
				Details: []string{},
			}
			details := v.Details()
			if len(details) > 0 {
				response.Details = details
			}
			ctx.JSON(v.StatusCode(), response)
			return
		}
	}

	ctx.JSON(http.StatusOK, Response{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: gin.H{},
	})
}
