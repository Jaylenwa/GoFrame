package util

import (
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

func NowTimeFormatRfc3339() {
	time.Now().Format("2006-01-02T15:04:05.999999-07:00")
}

func NowTimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowTimeFormatExecuteId() string {
	return gconv.String(time.Now().UnixNano())
}
