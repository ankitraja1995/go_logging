package go_logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"reflect"
	"strings"
)

type FluentdLogger struct {
	FluentdConnection *fluent.Fluent
	InitLogDetails    *LogConfig
	FileLogger        *logrus.Logger
}

func (fluentdLogger *FluentdLogger) Info( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Info, args...)
	}
}

func (fluentdLogger *FluentdLogger) Debug( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Debug, args...)
	}
}

func (fluentdLogger *FluentdLogger) Trace( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Trace, args...)
	}
}

func (fluentdLogger *FluentdLogger) Error( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Error,  args...)
	}
}

func (fluentdLogger *FluentdLogger) Fatal( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Fatal, args...)
	}
}

func (fluentdLogger *FluentdLogger) Panic( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Panic, args...)
	}
}

func (fluentdLogger *FluentdLogger) Warn( tag string, format string, args ...interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncLog(tag, format, Warn, args...)
	}
}

func (fluentdLogger *FluentdLogger) asyncLog( tag string,  format string,  logType LogType, args ...interface{}) {
	data := inputArgsToMap(format,args...)
	logLevel := GetLogLevelFromLogType(logType)
	loggerChain := GetChainOfLoggers()
	fluentdPostErr := loggerChain.Execute(fluentdLogger, logLevel, tag, data)

	if fluentdPostErr != nil {
		writeTofile(fluentdLogger, format, logType, false)
	}
}

func (fluentdLogger *FluentdLogger) EventLog(tag string, obj interface{}) {
	if fluentdLogger.InitLogDetails.IsEnabled {
		go fluentdLogger.asyncEventLog(tag, obj)
	}
}

func (fluentdLogger *FluentdLogger) asyncEventLog(tag string, obj interface{}) {
	bytesObj, jsonMarshallErr := json.Marshal(obj)
	if jsonMarshallErr != nil {
		jsonMarshallErr = Wrapf(jsonMarshallErr, "Error: while marshalling-- %v  for logging into file", obj)
		InternalLoggerGlobal.Error(jsonMarshallErr)
	}
	format :=  bytes.NewBuffer(bytesObj).String()
	data := inputArgsToMap(format)
	fluentdPostError := fluentdLogger.FluentdConnection.Post(tag, data)

	if fluentdPostError != nil {
		writeTofile(fluentdLogger, format, LogLevelForEvent , true)
	}
}

func GetLorusInstance(logFileConfig *LogFileConfig) *logrus.Logger {
	lg := lumberjack.Logger{
		Filename:   logFileConfig.LogFilename,
		MaxSize:    logFileConfig.MaxSize,
		MaxBackups: logFileConfig.MaxBackups,
		MaxAge:     logFileConfig.MaxAge,
		Compress:   logFileConfig.Compress,
	}
	Log := logrus.New()
	Log.SetOutput(&lg)
	logLevel, loglevelParseErr := logrus.ParseLevel(strings.ToLower(string(logFileConfig.GlobalLoggingType)))
	if loglevelParseErr != nil {
		loglevelParseErr = Wrapf(loglevelParseErr, "Error: while getting file level. kindly supply logging level from the "+
			"LogType enum only (trace, debug, info, warn, error, fatal, panic supported)")
		InternalLoggerGlobal.Error(loglevelParseErr)
	}
	Log.SetLevel(logLevel)

	return Log
}

func fileLogRouter(fileLogger *logrus.Logger, LogLevel string, data string) {
	t1 := fileLogger
	callLogrusFuncByName(t1, LogLevel, data)
}

func callLogrusFuncByName(logrusInterface *logrus.Logger, funcName string, params ...interface{}) {
	myClassValue := reflect.ValueOf(logrusInterface)
	m := myClassValue.MethodByName(funcName)
	if !m.IsValid() {
		fluentdPostError := FluentdPostError.Newf("Error: method not found-- %s, for logger-- %v, "+
			"having params-- %v", funcName, &logrusInterface, params)
		InternalLoggerGlobal.Error(fluentdPostError)
	} else {
		in := make([]reflect.Value, len(params))
		for i, param := range params {
			in[i] = reflect.ValueOf(param)
		}
		m.Call(in)
	}
}

func inputArgsToMap(format string, args ...interface{}) map[string]string {
	formatted_string := fmt.Sprintf(format, args...)
	data := map[string]string{
		CommonKeyForJsonData: formatted_string,
	}

	return data
}

func writeTofile(fluentdLogger *FluentdLogger, format string, logType LogType, isEvent bool) {
	jsonData, jsonMarshallErr := json.Marshal(format)
	if jsonMarshallErr != nil {
		jsonMarshallErr = Wrapf(jsonMarshallErr, "Error: while marshalling-- %v  for logging into file", format)
		InternalLoggerGlobal.Error(jsonMarshallErr)
	} else {
		fileLogRouter(fluentdLogger.FileLogger, string(logType), string(jsonData))
	}
}
