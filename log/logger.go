package log

import (
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
func Info(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Infof log
func Infof(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Print log
func Print(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Printf log
func Printf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Debug log
func Debug(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debug(args...)
}

// Debugf log
func Debugf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debugf(format, args...)
}

// Warn log
func Warn(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warn(args...)
}

// Warnf log
func Warnf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warnf(format, args...)
}

// Error log
func Error(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Error(args...)
}

// Errorf log
func Errorf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Errorf(format, args...)
}

// Fatal log
func Fatal(args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatal(args...)
}

// Fatalf log
func Fatalf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatalf(format, args...)
}
