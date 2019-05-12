package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfgFile string) error {
	cfg := Config{
		Name: cfgFile,
	}
	//初始化配置文件
	if err := cfg.initConfig(); err != nil {
		return err
	}

	//初始化日志配置
	cfg.initLog()

	//监控配置文件变化
	cfg.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	var (
		err error
	)
	//如果没有指定配置文件则使用默认配置
	if c.Name == "" {
		viper.AddConfigPath("conf")   //默认配置文件目录
		viper.SetConfigName("config") //默认配置文件名
	} else {
		viper.SetConfigFile(c.Name) //使用指定的配置文件
	}
	viper.SetConfigType("yaml")  //设置配置文件的格式yaml
	viper.AutomaticEnv()         //读取匹配的环境变量
	viper.SetEnvPrefix("SERVER") //读取环境变量前缀SERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	//viper解析配置文件
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("配置文件发生改变:%s", in.Name)
	})

}

//初始化日志文件
func (c *Config) initLog() {
	passLargerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.log_level"),
		LoggerFile:     viper.GetString("log.log_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rolling_policy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLargerCfg)
}
