package logger

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func Test_Log(t *testing.T) {
	New(zap.DebugLevel, "", "", 0)

	var loggerEntity = Get()
	loggerEntity.Debug("log 初始化成功")
	loggerEntity.Info("无法获取网址",

		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

}
