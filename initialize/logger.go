package initialize

import (
	"os"
	"path"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化日志 https://github.com/uber-go/zap
func InitLogger() {
	logConfig := global.Config.Log
	// 判断日志目录是否存在，不存在则创建
	if exist := utils.IsDirExists(logConfig.Path); !exist {
		_ = os.Mkdir(logConfig.Path, os.ModePerm)
	}
	// 设置输出格式
	var encoder zapcore.Encoder
	if logConfig.Format == "json" {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	// 设置文件归档
	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer())
	// 创建zap实例
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel())
	// 创建logger实例
	logger := zap.New(zapCore)
	defer logger.Sync()
	// 赋值给global
	global.Log = logger
	// Logger
	zap.ReplaceGlobals(global.Log)
}

// getEncoderConfig 自定义日志输出字段
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime, // 自定义输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return config
}

// getEncodeTime 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// 获取文件切割和归档配置信息
func getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberjackConfig := global.Config.Log.Archive
	// https://github.com/natefinch/lumberjack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   getLogFile(),                //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	// 设置日志文件切割
	return zapcore.AddSync(lumberjackLogger)
}

// 获取日志文件名
func getLogFile() string {
	return path.Join(global.Config.Log.Path, global.Config.Log.Level+".log")
}

// 获取日志级别
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
	if level, ok := levelMap[global.Config.Log.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}
