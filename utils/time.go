package utils

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func ParseTimeToTimeStamp(timeStr string) int64 {
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		log.Error(err)
	}
	return parse.UnixNano() / 1e6
}
