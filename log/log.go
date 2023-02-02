package log

import (
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const CallerSkipNum = 1

var (
	s *zap.SugaredLogger
)

func zapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stack",
		LineEnding:    "\n",
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, e zapcore.PrimitiveArrayEncoder) {
			e.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

}
func init() {

	_ = NewLoggerSugar("default", "", "debug")
}

func NewLoggerSugar(serviceName, logFile string, level string) error {

	levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel(level)
	})

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(zapEncoderConfig()), zapcore.AddSync(os.Stdout), levelEnabler))

	if logFile != "" {
		hook := &lumberjack.Logger{
			Filename:   logFile, // 日志文件路径
			MaxSize:    128,     // 每个日志文件保存的大小 单位:M
			MaxAge:     7,       // 文件最多保存多少天
			MaxBackups: 30,      // 日志文件最多保存多少个备份
			Compress:   false,   // 是否压缩
		}
		fileWriter := zapcore.AddSync(hook)
		writes := []zapcore.WriteSyncer{fileWriter}
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(zapEncoderConfig()),
			zapcore.NewMultiWriteSyncer(writes...),
			levelEnabler,
		))

	}

	tree := zapcore.NewTee(cores...)
	logger := zap.New(tree, zap.WithCaller(true), zap.AddCallerSkip(CallerSkipNum), zap.AddStacktrace(zapcore.ErrorLevel))
	logger.With(zap.String("service_name", serviceName))
	s = logger.Sugar()
	return nil
}

// LogLevel log level
func logLevel(level string) zapcore.Level {

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
func Debug(args ...interface{}) {
	s.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	s.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	s.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	s.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	s.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	s.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	s.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	s.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	s.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	s.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	s.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	s.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	s.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	s.Fatalf(template, args...)
}
