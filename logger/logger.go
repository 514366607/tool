package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	lock   sync.RWMutex
)

// Get 取得数据
func Get() *zap.Logger {
	lock.RLock()
	defer lock.RUnlock()
	return logger
}

// New 初始化
// 日志等级 debug info warn error fatal
func New(level zapcore.Level, serviceName, filePath string, serverID uint32) {
	lock.Lock()
	defer lock.Unlock()

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevelAt(level)

	var encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间格式化
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var ioWrite = make([]zapcore.WriteSyncer, 0)

	var options = []zap.Option{
		zap.AddCaller(), // 开启文件及行号
	}

	// 打印到控制台
	if level.Enabled(zapcore.DebugLevel) || filePath == "" {
		// 全路径编码器
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

		// 打印stack
		options = append(options, zap.AddStacktrace(zapcore.DebugLevel))

		// 开启开发模式，堆栈跟踪
		options = append(options, zap.Development())

		ioWrite = append(ioWrite, zapcore.AddSync(os.Stdout))
	}
	if filePath != "" {
		// 打印到文件
		hook := lumberjack.Logger{
			Filename:   filePath, // 日志文件路径
			MaxSize:    128,      // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,       // 日志文件最多保存多少个备份
			MaxAge:     7,        // 文件最多保存多少天
			Compress:   true,     // 是否压缩
		}
		ioWrite = append(ioWrite, zapcore.AddSync(&hook))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),   // 编码器配置
		zapcore.NewMultiWriteSyncer(ioWrite...), // 打印到那里
		atomicLevel,                             // 日志级别
	)

	// 设置初始化字段
	options = append(options, zap.Fields(zap.String("serviceName", serviceName)))
	options = append(options, zap.Fields(zap.Uint32("server_id", serverID)))
	// 构造日志
	logger = zap.New(core, options...)
}
