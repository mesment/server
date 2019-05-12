package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/model"
	"github.com/mesment/server/pkg/errno"
	"github.com/mesment/server/pkg/token"
)

func Login(c *gin.Context) {
	var user model.UserModel
	var d *model.UserModel
	var err error

	//绑定用户数据到用户结构体
	if err := c.Bind(&user); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//查询用户信息
	if d, err = model.GetUser(user.Username); err != nil {
		handler.SendResponse(c, errno.ErrUserNotExist, nil) //用户不存在
		return
	}

	//比较用户密码是否正确
	if d.Password != user.MD5Password() {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil) //密码错误
		return
	}

	tokenStr, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil) //token错误
		return
	}

	//返回token
	handler.SendResponse(c, nil, model.Token{Token: tokenStr})
}
