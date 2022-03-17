package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	c "github.com/skrbox/ioseek/pkg/conf"
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

	once sync.Once
	L    *zap.SugaredLogger // 全局 logger
)

func init() {
	once.Do(func() {
		var encoder zapcore.Encoder
		switch *c.LogStyle {
		case c.Json:
			encoder = zapcore.NewJSONEncoder(config)
		case c.Txt:
			encoder = zapcore.NewConsoleEncoder(config)
		default:
			encoder = zapcore.NewJSONEncoder(config)
		}
		logger := zap.New(zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.ErrorLevel),
		), zap.AddCaller())
		defer func() { _ = logger.Sync() }()
		L = logger.Sugar()
	})
}
