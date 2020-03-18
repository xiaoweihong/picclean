package contorller

import (
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"picclean/utils"
)

func DbDeleteResult(engine *xorm.Engine) {
	isDeleteDB := viper.GetBool("garbage.deleteDB")
	if isDeleteDB {
		log.Infof("开始删除数据库记录，时间区间为[%v %v]", utils.ParseTimeStampToTime(startTime), utils.ParseTimeStampToTime(endTime))
		DeleteResult(engine)
	} else {
		log.Infof("不删除数据库记录，已经删除指定时间区间[%v  %v]的图片", utils.ParseTimeStampToTime(startTime), utils.ParseTimeStampToTime(endTime))
	}
}
