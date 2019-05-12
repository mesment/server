package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var DB *Database

type Database struct {
	DB     *gorm.DB
	Docker *gorm.DB
}

func (db *Database) Init() {
	DB = &Database{
		DB:     GetDB(),
		Docker: GetDockerDB(),
	}
}

func openDB(username, password, addr, dbname string) *gorm.DB {
	cfg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s",
		username,
		password,
		addr,
		dbname,
		"Local",
	)

	fmt.Println(cfg)
	db, err := gorm.Open("mysql", cfg)
	if err != nil {
		log.Errorf(err, "连接数据库失败.dbname:%s", dbname)
	}

	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	//用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(0)

	//设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，
	// 可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(2000)
}

//关闭数据库
func (db *Database) Close() {
	db.DB.Close()
	db.Docker.Close()
}

func InitDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.db_name"))
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.db_name"))
}

func GetDB() *gorm.DB {

	return InitDB()
}

func GetDockerDB() *gorm.DB {

	return InitDockerDB()
}
