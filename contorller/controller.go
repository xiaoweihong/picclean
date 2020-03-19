package contorller

import (
	"fmt"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"picclean/entity"
	"picclean/utils"
	"strings"
	"sync"
)

var (
	startTime     int64
	endTime       int64
	tables        []string
	tablesNoIndex int
	picCounts     int64
	imageChan     = make(chan string, viper.GetInt("garbage.maxChannel"))
	tableCount    int
	wg            sync.WaitGroup
)

// QueryResult  根据时间区间(startTime,endTime)获取table的图片地址
func QueryResult(table string, engine *xorm.Engine, q *entity.QueryTime) {
	_ = engine.
		Where("ts>?", q.StartTime).
		Where("ts<?", q.EndTime).
		Table(table).Iterate(new(entity.ImageURL), func(idx int, bean interface{}) error {
		imageURL := bean.(*entity.ImageURL)
		if imageURL.ImageURI != "" || imageURL.CutboardImageURI != "" {
			imageChan <- utils.ConverArceeURLToWeedUrl(imageURL.ImageURI)
			imageChan <- utils.ConverArceeURLToWeedUrl(imageURL.CutboardImageURI)
		}
		return nil
	})
	tableCount++
	if tableCount == tablesNoIndex {
		close(imageChan)
	}
}

// GetAllResult 获得所有表中的图片url数据切片
func GetAllResult(engine *xorm.Engine) {
	startTime = utils.ParseTimeToTimeStamp(viper.GetString("deletetime.startTime"))
	endTime = utils.ParseTimeToTimeStamp(viper.GetString("deletetime.endTime"))
	tables = viper.GetStringSlice("postgres.tables")
	queryTime := &entity.QueryTime{
		StartTime: startTime,
		EndTime:   endTime,
	}

	for _, table := range tables {
		if !strings.Contains(table, "_index") {
			tablesNoIndex++
			go QueryResult(table, engine, queryTime)
		}
	}
	//log.WithFields(log.Fields{"toal": len(allResultSlice)}).Info(allResultSlice)
}

// DeleteResultFromDB 根据时间区间(startTime,endTime)删除table的记录
func DeleteResultFromDB(table string, engine *xorm.Engine, q *entity.QueryTime) {
	count, err := engine.
		Where("ts>?", q.StartTime).
		Where("ts<?", q.EndTime).
		Table(table).
		Delete(new(entity.ImageURL))
	log.WithFields(log.Fields{
		"数据表": table,
		"数量":  count,
	}).Info("数据删除统计")
	if err != nil {
		log.Error(err)
	}
}

// DeleteResult 删除table的记录
func DeleteResult(engine *xorm.Engine) {
	startTime = utils.ParseTimeToTimeStamp(viper.GetString("deletetime.startTime"))
	endTime = utils.ParseTimeToTimeStamp(viper.GetString("deletetime.endTime"))
	tables = viper.GetStringSlice("postgres.tables")
	queryTime := &entity.QueryTime{
		StartTime: startTime,
		EndTime:   endTime,
	}
	for _, table := range tables {
		DeleteResultFromDB(table, engine, queryTime)
	}
}

// DeleteUrl 根据weed的api删除图片
func DeleteUrlFromWeed() {
	for img := range imageChan {
		log.Debug("delete URL-->", img)
		err := utils.Delete(img, "")
		if err != nil {
			fmt.Println(err)
		}
		picCounts++
	}
	wg.Done()
}

// 删除图片
func DelURL(engine *xorm.Engine) {
	GetAllResult(engine)
	parallelNum := viper.GetInt("garbage.parallelNum")
	for i := 0; i < parallelNum; i++ {
		wg.Add(1)
		DeleteUrlFromWeed()
	}
	wg.Wait()
}

func CountAndGarbage() {
	garbageURL := viper.GetString("garbage.URL")
	garbageThreshold := viper.GetFloat64("garbage.garbageThreshold")
	reqURL := fmt.Sprintf("http://%s/vol/vacuum?garbageThreshold=%f", garbageURL, garbageThreshold)
	log.WithField("totalCount", picCounts).Info("总计删除图片数据")
	log.WithField("reqURL", reqURL).Info("开始垃圾回收释放空间")
	get, err := utils.Get(reqURL)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(string(get))
}
