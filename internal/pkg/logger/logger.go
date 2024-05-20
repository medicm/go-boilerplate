package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	*logrus.Logger
}

var (
	logger     *logrus.Logger
	onceLogger sync.Once
)

func GetLoggerInstance() *logrus.Logger {
	onceLogger.Do(func() {
		logger = logrus.New()
		logger.Out = os.Stdout
		logger.Formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		}

		levelStr := viper.GetString("APP_LOG_LEVEL")
		logger.Println("Log level is", levelStr)
		if level, err := logrus.ParseLevel(strings.ToUpper(levelStr)); err == nil {
			logger.Println("Setting log level to", levelStr)
			logger.Level = level
		} else {
			logger.Level = logrus.InfoLevel
		}
	})
	return logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: GetLoggerInstance(),
	}
}
