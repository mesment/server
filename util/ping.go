package util

import (
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

//检查服务器是否正常
func Ping() error {
	var (
		resp *http.Response
		err  error
	)

	for i := 0; i < viper.GetInt("max_ping_count"); i++ {

		resp, err = http.Get(viper.GetString("url") + "/health/check")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		//检查失败，睡眠1秒后重新尝试
		log.Info("Ping服务器失败,1秒后重新尝试")
		//sleep 1 second
		time.Sleep(time.Second)
	}

	return errors.New("连接服务器失败")

}
