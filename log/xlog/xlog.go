package xlog

import (
	"context"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var sugaredLogger *zap.SugaredLogger

const (
	serviceNameKey = "service_name"
)

func init() {
	_ = InitXLog(zapcore.DebugLevel, "console_service")
}

func InitXLog(logLevel zapcore.Level, serviceName string) error {

	config := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	jsonEncoder := zapcore.NewJSONEncoder(config)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel
	})

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), infoLevel))

	c := zapcore.NewTee(cores...)
	fields := make([]zap.Field, 0)
	fields = append(fields, zap.String(serviceNameKey, serviceName))
	logger = zap.New(c, zap.WithCaller(true), zap.AddCallerSkip(1), zap.Fields(fields...))
	sugaredLogger = logger.Sugar()
	return nil
}

// log level
func LogLevel(level string) zapcore.Level {

	level = strings.ToLower(level)
	switch level {

	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(ctx context.Context, args ...interface{}) {

	sugaredLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(ctx context.Context, args ...interface{}) {

	sugaredLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(ctx context.Context, args ...interface{}) {

	sugaredLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(ctx context.Context, args ...interface{}) {

	sugaredLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(ctx context.Context, args ...interface{}) {

	sugaredLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(ctx context.Context, args ...interface{}) {

	sugaredLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(ctx context.Context, args ...interface{}) {

	sugaredLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(ctx context.Context, template string, args ...interface{}) {

	sugaredLogger.Fatalf(template, args...)
}
