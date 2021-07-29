package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/pkg/errcode"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: data,
	})
}

func (r *Response) Error(c *gin.Context, err *errcode.Error, args ...interface{}) {
	response := gin.H{"code": err.Code(), "msg": fmt.Sprintf(err.Msg(), args...), "data": gin.H{}}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	c.JSON(err.StatusCode(), response)
}
