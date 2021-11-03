package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	timeLayout = "2006-01-02"
	AppName    = "SM"
)

var (
	log *logrus.Logger
)

func Init(fileName string) {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.SetLevel(logrus.PanicLevel)
	log.SetLevel(logrus.FatalLevel)
	log.SetLevel(logrus.ErrorLevel)
	log.SetLevel(logrus.WarnLevel)
	log.SetLevel(logrus.InfoLevel)
	log.SetLevel(logrus.DebugLevel)
	log.SetLevel(logrus.InfoLevel)

	if fileName == "" {
		log.Out = os.Stdout
		return
	}

	file, err := os.OpenFile(fmt.Sprintf(fileName, time.Now().Format(timeLayout)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr, ", err)
	}
}

func getFileAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		i := strings.Index(file, AppName)
		if i >= 0 {
			file = file[i:]
		}
	}

	return file, line
}

// Info log
func Info(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Info(args...)
}

// Infof log
func Infof(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Infof(format, args...)
}

// Print log
func Print(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Info(args...)
}

// Printf log
func Printf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Infof(format, args...)
}

// Debug log
func Debug(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Debug(args...)
}

// Debugf log
func Debugf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Debugf(format, args...)
}

// Warn log
func Warn(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Warn(args...)
}

// Warnf log
func Warnf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Warnf(format, args...)
}

// Error log
func Error(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Error(args...)
}

// Errorf log
func Errorf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Errorf(format, args...)
}

// Fatal log
func Fatal(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Fatal(args...)
}

// Fatalf log
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Fatalf(format, args...)
}
