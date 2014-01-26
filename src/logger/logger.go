//	日志

package logger

import (
	"fmt"
	"os"
)

//	日志等级
const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelCritical
)

// 日志等级
var logLevel = LogLevelDebug

// 获取日志等级
func GetLevel() int {
	return logLevel
}

// 设置日志等级
func SetLevel(l int) {
	logLevel = l
}

// 日志时间格式
var timeFormat = "2013-04-02 10:36:22 Z07:00"

func SetFormat(f string) {
	timeFormat = f
}

// 日志接口
type Logger interface {
	Debug(info, message string)
	Info(info, message string)
	Warning(info, message string)
	Error(info, message string)
	Critical(info, message string)
}

// 日志
var logger = NewStandardLogger(os.Stdout)

func SetLogger(l Logger) {
	logger = l
}

func Debugf(format string, args ...interface{}) {
	if logLevel <= LogLevelDebug {
		ci := retrieveCallInfo()
		fi := fmt.Sprintf(format, args...)

		logger.Debug(ci.verboseFormat(), fi)
	}
}

func Infof(format string, args ...interface{}) {
	if logLevel <= LogLevelInfo {
		ci := retrieveCallInfo()
		fi := fmt.Sprintf(format, args...)

		logger.Info(ci.verboseFormat(), fi)
	}
}

func Warningf(format string, args ...interface{}) {
	if logLevel <= LogLevelWarning {
		ci := retrieveCallInfo()
		fi := fmt.Sprintf(format, args...)

		logger.Warning(ci.verboseFormat(), fi)
	}
}

func Errorf(format string, args ...interface{}) {
	if logLevel <= LogLevelError {
		ci := retrieveCallInfo()
		fi := fmt.Sprintf(format, args...)

		logger.Error(ci.verboseFormat(), fi)
	}

}

func Criticalf(format string, args ...interface{}) {
	if logLevel <= LogLevelCritical {
		ci := retrieveCallInfo()
		fi := fmt.Sprintf(format, args...)

		logger.Critical(ci.verboseFormat(), fi)
	}
}
