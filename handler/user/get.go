package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/model"
	"github.com/mesment/server/pkg/errno"
)

func Get(c *gin.Context) {
	var user *model.UserModel
	var err error
	name := c.Param("username")

	if user, err = model.GetUser(name); err != nil {
		handler.SendResponse(c, errno.ErrUserNotExist, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
