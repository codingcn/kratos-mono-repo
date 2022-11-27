package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func NewLogger(isProd bool, logLevel zapcore.Level, logFileName string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logFileName, // 日志文件路径
		MaxSize:    1024,        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10,          // 日志文件最多保存多少个备份
		MaxAge:     14,          // 文件最多保存多少天
		Compress:   true,        // 是否压缩
	}
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	if isProd {
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "Logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
				EncodeTime:     timeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.FullCallerEncoder,
			}), // 编码器配置
			//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),      // 打印到控制台
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
			logLevel, // 日志级别
		)

		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// 开启文件及行号
		development := zap.Development()
		return zap.New(core, caller, development, zap.AddStacktrace(zap.ErrorLevel), zap.AddCallerSkip(3))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		}), // 编码器配置
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),      // 打印到控制台
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		logLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	return zap.New(core, caller, zap.AddStacktrace(zap.ErrorLevel), zap.AddCallerSkip(3))

}
