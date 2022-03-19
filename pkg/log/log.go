package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	config = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder = zapcore.NewJSONEncoder(config)

	once sync.Once
	L    *zap.SugaredLogger // 全局 logger
	CL   *cronLogger        // 定时任务logger
)

func init() {
	once.Do(func() {
		logger := zap.New(zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.ErrorLevel),
		), zap.AddCaller())
		defer func() { _ = logger.Sync() }()
		L = logger.Sugar()
		CL = &cronLogger{L}
	})
}
