package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func init() {
	log = logrus.New()

	// Ensure log directory exists
	if err := os.MkdirAll("./log", os.ModePerm); err != nil {
		log.Fatal("Failed to create log directory: ", err)
	}

	// Configure lumberjack for log rotation
	rotateFileHook := &lumberjack.Logger{
		Filename:   "./log/app.log", // Changed to use relative path
		MaxSize:    100,             // Max size in megabytes before rotation
		MaxBackups: 3,               // Max number of old log files to retain
		MaxAge:     28,              // Max number of days to retain old log files
		Compress:   true,            // Compress rotated files
	}

	// Set logrus to use lumberjack for output
	log.SetOutput(rotateFileHook)

	// Optional: Set log level and format
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
}

// Export log functions for use in other packages
func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
