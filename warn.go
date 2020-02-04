package go_logging

type warn struct {
	next  abstractLogger
	level LogLevel
}

func (wrn *warn) setNext(next abstractLogger) {
	wrn.next = next
}

func (wrn *warn) setLevel(level LogLevel) {
	wrn.level = level
}

func (wrn *warn) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if wrn.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if wrn.next != nil {
		chainErr := wrn.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
