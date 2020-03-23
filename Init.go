package go_logging

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type LogConfig struct {
	IsEnabled                bool
	FluentDHost              string
	FluentDPort              int
	LogFilename              string
	MaxBackups               int
	MaxSize                  int
	MaxAge                   int
	Compress                 bool
	GlobalLoggingType        LogType
	FluentdPostTimeoutMillis time.Duration
	InternalLogPath          string
	MaxRetry                 int
	MaxRetryWaitMillis       int
}

type LogFileConfig struct {
	LogFilename       string
	MaxBackups        int
	MaxSize           int
	MaxAge            int
	Compress          bool
	GlobalLoggingType LogType
}

func (lc *LogConfig) GetLogger() *FluentdLogger {
	var fluentdLogger FluentdLogger
	setInternalLogger(lc)

	logger, fluentDConnectionErr := fluent.New(fluent.Config{FluentPort: lc.FluentDPort, FluentHost: lc.FluentDHost, MaxRetry: lc.MaxRetry, MaxRetryWait: lc.MaxRetryWaitMillis})
	if fluentDConnectionErr != nil {
		fluentDConnectionErr = Wrapf(fluentDConnectionErr, "\n Error: could not create a new fluentd logger. Logging to file instead of fluentd.")
		InternalLoggerGlobal.Error(fluentDConnectionErr)
	}
	defer logger.Close()
	logFileConfig := LogFileConfig{
		LogFilename:       lc.LogFilename,
		MaxBackups:        lc.MaxBackups,
		MaxSize:           lc.MaxSize,
		MaxAge:            lc.MaxAge,
		Compress:          lc.Compress,
		GlobalLoggingType: lc.GlobalLoggingType,
	}

	fileLogger := GetLorusInstance(&logFileConfig)
	fileLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "data",
		},
			})
	fluentdLogger = FluentdLogger{
		FluentdConnection: logger,
		InitLogDetails:    lc,
		FileLogger:        fileLogger,
	}

	return &fluentdLogger
}

var logConfig sync.Once
var InternalLoggerGlobal *logrus.Logger

func setInternalLogger(lg *LogConfig) {
	logConfig.Do(func() {
		logFileConfig := LogFileConfig{
			LogFilename:       lg.InternalLogPath,
			MaxBackups:        lg.MaxBackups,
			MaxSize:           lg.MaxSize,
			MaxAge:            lg.MaxAge,
			Compress:          lg.Compress,
			GlobalLoggingType: lg.GlobalLoggingType,
		}
		InternalLoggerGlobal = GetLorusInstance(&logFileConfig)
	})
}
