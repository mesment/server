package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/server/pkg/errno"

	"net/http"
)

//返回信息
type Respone struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Respone{
		Code: code,
		Msg:  msg,
		Data: data,
	})

}
