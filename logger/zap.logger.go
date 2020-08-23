package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type logger struct {
	basePath string
	suffix string
	sugaredLogger *zap.SugaredLogger
}

var split = "."
var suffix = split + "log"
var logPath = "/code/log/"
var l logger

/**
将日志写入文件而不是终端
zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
 */
func InitZapLogger(serverName string) error {
	l = logger{
		basePath: logPath,
	}

	l.suffix = suffix
	logPrefix := getLogPrefix()
	fileName := l.basePath + logPrefix + split + serverName + l.suffix

	encoder := zapcore.NewJSONEncoder(initEncoderConfig())

	fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(fp), zapcore.DebugLevel)
	// 添加调用函数信息记录到日志中
	logger := zap.New(core, zap.AddCaller())
	l.sugaredLogger = logger.Sugar()

	return nil
}

func getLogPrefix() string {
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()

	return fmt.Sprintf("%d%d%d", year, month, day)
}

func initEncoderConfig() zapcore.EncoderConfig  {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.RFC3339TimeEncoder
	conf.EncodeLevel = zapcore.CapitalLevelEncoder
	return conf
}

func Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	l.sugaredLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}
