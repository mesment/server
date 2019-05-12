package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//健康检查返回OK
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
