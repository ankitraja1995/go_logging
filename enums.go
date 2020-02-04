package go_logging

type LogType string

const (
	Trace = LogType("Trace")
	Debug = LogType("Debug")
	Info  = LogType("Info")
	Warn  = LogType("Warn")
	Error = LogType("Error")
	Fatal = LogType("Fatal")
	Panic = LogType("Panic")
)

type LogLevel int

const (
	Trace_level = LogLevel(1)
	DebugLevel  = LogLevel(2)
	InfoLevel   = LogLevel(3)
	WarnLevel   = LogLevel(4)
	ErrorLevel  = LogLevel(5)
	FatalLevel  = LogLevel(6)
	PanicLevel  = LogLevel(7)
)

var LogLevelFromTypeMap = map[string]LogLevel{
	"TraceLevel": Trace_level,
	"DebugLevel": DebugLevel,
	"InfoLevel":  InfoLevel,
	"WarnLevel":  WarnLevel,
	"ErrorLevel": ErrorLevel,
	"FatalLevel": FatalLevel,
	"PanicLevel": PanicLevel,
}

func GetLogLevelFromLogType(logType LogType) LogLevel {
	logging_level := string(logType) + "Level"
	passedLogLevel := LogLevelFromTypeMap[logging_level]
	return passedLogLevel
}
