package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/model"
	"github.com/mesment/server/pkg/errno"
	"strconv"
)

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	log.Infof("DELETE ID:%d", id)
	if err := model.DeleteUser(uint64(id)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
