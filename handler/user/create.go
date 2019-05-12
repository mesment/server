package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/model"
	"github.com/mesment/server/pkg/errno"
)

func Create(c *gin.Context) {
	var user CreateRequest
	var resp CreateResponse

	if err := c.Bind(&user); err != nil {
		//返回响应信息
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	/*
		admin := c.Param("username")
		log.Infof("URL username:%s", admin)

		desc := c.Query("desc")
		log.Infof("URL key param desc:%s", desc)

		contentType := c.GetHeader("Content-Type")
		log.Infof("Content-Type:%s", contentType)

		log.Debugf("username:%s, password:%s", user.Username, user.Password)
	*/

	log.Infof("username:%s,passwd:%s", user.Username, user.Password)
	u := model.UserModel{
		Username: user.Username,
		Password: user.Password,
	}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	//判断用户名是否已存在，存在则报错
	if exist := model.IsUsernameExist(u.Username); exist {
		handler.SendResponse(c, errno.ErrUserAlreadyExist, nil)
		return
	}

	//加密用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	//插入用户信息到数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp = CreateResponse{
		Username: user.Username,
	}

	handler.SendResponse(c, nil, resp)

}
