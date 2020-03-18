package main

import (
	"flag"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"picclean/config"
	"picclean/contorller"
	"picclean/db"
	"sync"
	"time"
)

var (
	engine *xorm.Engine
	wg     *sync.WaitGroup
)

func init() {
	log.SetReportCaller(false)
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.0000",
	})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	// 解析配置文件
	workDir, _ := os.Getwd()
	config.InitConfig(workDir)
	engine = db.GetDBEngine()

	// 设置日志级别
	isDebug := viper.GetBool("log.debug")
	if isDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	flag.Parse()
}

func main() {
	sT := time.Now()
	contorller.DelURL(engine)
	defer engine.Close()

	log.Info("开始垃圾回收")
	contorller.CountAndGarbage()

	contorller.DbDeleteResult(engine)

	log.WithFields(log.Fields{
		"cost ": time.Since(sT),
	}).Info("共花费时间")
}
