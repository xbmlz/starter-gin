package logger

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/xbmlz/starter-gin/core/config"
	"github.com/xbmlz/starter-gin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Setup global log, see https://github.com/uber-go/zap
func Setup() {
	fmt.Println(config.App.Log.Path)
	if ok := utils.IsDirExists(config.App.Log.Path); !ok {
		fmt.Printf("create %v directory\n", config.App.Log.Path)
		_ = os.Mkdir(config.App.Log.Path, os.ModePerm)
	}
	// 设置输出格式
	encoder := zapcore.NewJSONEncoder(getEncoderConfig())
	// 设置日志文件切割
	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel())
	// 创建logger
	logger := zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	defer logger.Sync()
	// 赋值给全局变量
	// global.GvaLogger = logger
}

// 获取最低记录日志级别
func getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[config.App.Log.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}

// 自定义日志输出字段
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// 获取文件切割和归档配置信息
func getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename: getLogFile(), //日志文件
		// MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		// MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		// MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		// Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

// 获取日志文件名
func getLogFile() string {
	fileFormat := time.Now().Format("global.GvaConfig.Log.FileFormat")
	fileName := strings.Join([]string{fileFormat, "log"}, ".")
	return path.Join(config.App.Log.Path, fileName)
}
