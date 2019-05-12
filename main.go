package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/lexkong/log"
	"github.com/mesment/server/config"
	"github.com/mesment/server/model"
	v "github.com/mesment/server/pkg/version"
	"github.com/mesment/server/router"
	"github.com/mesment/server/router/middleware"
	"github.com/mesment/server/util"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	version = pflag.BoolP("version", "v", false, "显示版本信息")
	cfg     = pflag.StringP("config", "c", "", "配置文件路径")
)

func main() {
	pflag.Parse()

	if *version {
		v := v.Get()
		marshaled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshaled))
		return
	}

	if err := config.Init(*cfg); err != nil {
		log.Infof("初始化配置文件失败:%s", err)
		panic(err)
	}
	//从配置文件读取服务器地址端口号
	host := viper.GetString("server_addr")

	//初始化数据库连接
	model.DB.Init()
	defer model.DB.Close()

	r := gin.New()

	//
	middlewares := []gin.HandlerFunc{middleware.Logging(), middleware.RequestId()}

	router.AddMiddleWare(r, middlewares...)

	log.Infof("开始监听:%s", host)

	//启一个协程检查服务器启动是否正常
	go func() {
		if err := util.Ping(); err != nil {
			log.Fatal("服务器响应异常:", err)
		}
		log.Info("服务器已正常启动")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")

	//如果有配置TLS证书和私钥则启用https
	if cert != "" && key != "" {
		log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
		log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, r).Error())
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("server_addr"))
	http.ListenAndServe(host, r)
	//r.Run(host)

}
