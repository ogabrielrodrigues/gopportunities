package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	logs_file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log_config := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "message",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	json_encoder := zapcore.NewJSONEncoder(log_config)

	logger = zap.New(
		zapcore.NewTee(
			zapcore.NewCore(json_encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
			zapcore.NewCore(json_encoder, zapcore.AddSync(logs_file), zapcore.InfoLevel),
		),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

func Info(msg string, tags ...zap.Field) {
	logger.Info(msg, tags...)
	logger.Sync()
}

func Err(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Error(msg, tags...)
	logger.Sync()
}
