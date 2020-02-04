package main


type trace struct {
	next  abstractLogger
	level LogLevel
}

func (trc *trace) setNext(next abstractLogger) {
	trc.next = next
}

func (trc *trace) setLevel(level LogLevel) {
	trc.level = level
}

func (trc *trace) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if trc.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if trc.next != nil {
		chainErr := trc.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
