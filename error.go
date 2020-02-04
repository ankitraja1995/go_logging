package go_logging

type errs struct {
	next  abstractLogger
	level LogLevel
}

func (er *errs) setNext(next abstractLogger) {
	er.next = next
}

func (er *errs) setLevel(level LogLevel) {
	er.level = level
}

func (er *errs) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if er.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if er.next != nil {
		chainErr := er.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
