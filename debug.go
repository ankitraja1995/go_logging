package go_logging

type debug struct {
	next  abstractLogger
	level LogLevel
}

func (dbg *debug) setNext(next abstractLogger) {
	dbg.next = next
}

func (dbg *debug) setLevel(level LogLevel) {
	dbg.level = level
}

func (dbg *debug) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if dbg.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if dbg.next != nil {
		chainErr := dbg.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
