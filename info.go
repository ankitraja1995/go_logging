package main

type info struct {
	next  abstractLogger
	level LogLevel
}

func (inf *info) setNext(next abstractLogger) {
	inf.next = next
}

func (inf *info) setLevel(level LogLevel) {
	inf.level = level
}

func (inf *info) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if inf.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if inf.next != nil {
		chainErr := inf.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
