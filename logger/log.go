package logger

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	// PanicLevel 异常
	PanicLevel int = iota
	// FatalLevel 语法错误
	FatalLevel
	// ErrorLevel 错误
	ErrorLevel
	// WarnLevel 警告
	WarnLevel
	// InfoLevel 信息
	InfoLevel
	// DebugLevel 调试
	DebugLevel
)

// LogFile 日志结构
type LogFile struct {
	level    int
	logTime  int64
	fileName string
	fileFd   *os.File
}

var logFile LogFile

// Config 配置
func Config(logFolder string, level int) {
	logFile.fileName = logFolder
	logFile.level = level

	//小于等于5才输出到文件
	if level <= 5 {
		log.SetOutput(logFile)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

// SetLevel 设置打印等级
func SetLevel(level int) {
	Config(logFile.fileName, level)
}

// GetLevel 取得Debug等级
func GetLevel() int {
	return logFile.level
}

// Debugf Debug打印
func Debugf(format string, args ...interface{}) {
	if logFile.level >= DebugLevel {
		log.SetPrefix("debug ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Debug Debug打印
func Debug(args ...interface{}) {
	if logFile.level >= DebugLevel {
		log.SetPrefix("debug ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

// Infof 信息打印
func Infof(format string, args ...interface{}) {
	if logFile.level >= InfoLevel {
		log.SetPrefix("info ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Info 信息打印
func Info(args ...interface{}) {
	if logFile.level >= InfoLevel {
		log.SetPrefix("info ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

// Warnf 警告打印
func Warnf(format string, args ...interface{}) {
	if logFile.level >= WarnLevel {
		log.SetPrefix("warn ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Warn 警告打印
func Warn(args ...interface{}) {
	if logFile.level >= WarnLevel {
		log.SetPrefix("warn ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

// Errorf 错误打印
func Errorf(format string, args ...interface{}) {
	if logFile.level >= ErrorLevel {
		log.SetPrefix("error ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Error 错误打印
func Error(args ...interface{}) {
	if logFile.level >= ErrorLevel {
		log.SetPrefix("error ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

// Fatalf 语法错误打印
func Fatalf(format string, args ...interface{}) {
	if logFile.level >= FatalLevel {
		log.SetPrefix("fatal ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Fatal 语法错误打印
func Fatal(args ...interface{}) {
	if logFile.level >= FatalLevel {
		log.SetPrefix("fatal ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

func (me LogFile) Write(buf []byte) (n int, err error) {
	if me.fileName == "" {
		fmt.Printf("consol: %s", buf)
		return len(buf), nil
	}

	if logFile.logTime+86400 < time.Now().Unix() {
		logFile.createLogFile()
		logFile.logTime = time.Now().Unix()
	}

	if logFile.fileFd == nil {
		return len(buf), nil
	}

	return logFile.fileFd.Write(buf)
}

func (me *LogFile) createLogFile() {
	logdir := "./"
	if index := strings.LastIndex(me.fileName, "/"); index != -1 {
		logdir = me.fileName[0:index] + "/"
		os.MkdirAll(me.fileName[0:index], os.ModePerm)
	}

	now := time.Now()
	filename := fmt.Sprintf("%s_%04d%02d%02d_%02d%02d", me.fileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	if err := os.Rename(me.fileName, filename); err == nil {
		go func() {
			tarCmd := exec.Command("tar", "-zcf", filename+".tar.gz", filename, "--remove-files")
			tarCmd.Run()

			//mac情况下会用不了删除命令，这里再删除一次。
			os.Remove(filename)

			rmCmd := exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +2 -exec rm {} \;`)
			rmCmd.Run()
		}()
	}

	fd, err := os.Create(me.fileName)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	me.fileFd = fd

	// for index := 0; index < 10; index++ {
	// 	if fd, err := os.OpenFile(me.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive); nil == err {
	// 		me.fileFd.Sync()
	// 		me.fileFd.Close()
	// 		me.fileFd = fd
	// 		break
	// 	}

	// 	me.fileFd = nil
	// }
}
