package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

type rotatingFileWriter struct {
	filename string
	size     int64
	maxSize  int64
}

// Init initializes the logger
func Init(logPath string, level string) {
	once.Do(func() {
		initLogger(logPath, level)
	})
}

func initLogger(logPath string, level string) {
	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(logPath), 0744); err != nil {
		panic(err)
	}

	// 设置日志级别
	zapLevel := getZapLevel(level)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 创建一个基于时间的日志文件名
	timeBasedLogPath := getTimeBasedLogPath(logPath)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&rotatingFileWriter{filename: timeBasedLogPath, maxSize: 10 * 1024 * 1024}), // 10MB
		),
		zapLevel,
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

func getTimeBasedLogPath(basePath string) string {
	dir := filepath.Dir(basePath)
	ext := filepath.Ext(basePath)
	base := filepath.Base(basePath[:len(basePath)-len(ext)])
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	return filepath.Join(dir, fmt.Sprintf("%s_%s%s", base, timestamp, ext))
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func (w *rotatingFileWriter) Write(p []byte) (n int, err error) {
	if w.size+int64(len(p)) >= w.maxSize {
		// 创建新的日志文件，而不是覆盖旧文件
		w.filename = getTimeBasedLogPath(w.filename)
		w.size = 0
	}
	f, err := os.OpenFile(w.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err = f.Write(p)
	w.size += int64(n)
	return n, err
}

// Debug logs a message at DebugLevel
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel and then calls os.Exit(1)
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
