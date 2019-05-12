package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/handler/user"
	"github.com/mesment/server/router/middleware"
	"github.com/mesment/server/router/middleware/auth"
	"net/http"
)

//加载中间件
func AddMiddleWare(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//处理404
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "Wrong API.")
	})

	g.POST("/login", user.Login)

	health := g.Group("/health")
	{
		health.GET("/check", handler.HealthCheck)
	}

	u := g.Group("/v1/user")
	//添加分组中间件
	u.Use(auth.AuthMiddleware())
	{
		u.POST("", user.Create)       //创建用户
		u.DELETE("/:id", user.Delete) //删除用户
		u.GET("/:username", user.Get) //查询用户详细信息
		u.PUT("/:id", user.Update)    //更新用户
		u.GET("", user.List)          //查看用户列表
	}

	return g

}
