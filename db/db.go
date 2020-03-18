package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetDBEngine() *xorm.Engine {
	// 数据库配置连接相关
	host := viper.GetString("postgres.host")
	port := viper.GetInt("postgres.port")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")

	// 数据库参数相关
	showsql := viper.GetBool("postgres.showSQL")
	showtime := viper.GetBool("postgres.showtime")
	maxLifeTime := viper.GetDuration("postgres.connMaxLifetime")
	maxOpenConns := viper.GetInt("postgres.maxOpenConns")
	maxIdleConns := viper.GetInt("postgres.maxIdleConns")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL(showsql) //菜鸟必备
	engine.ShowExecTime(showtime)
	engine.SetConnMaxLifetime(maxLifeTime)
	engine.SetMaxOpenConns(maxOpenConns)
	engine.SetMaxIdleConns(maxIdleConns)

	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return engine
}
