package main

type panic struct {
	next  abstractLogger
	level LogLevel
}

func (pnic *panic) setNext(next abstractLogger) {
	pnic.next = next
}

func (pnic *panic) setLevel(level LogLevel) {
	pnic.level = level
}

func (pnic *panic) Execute(fluentdLogger *FluentdLogger, passedLogLevel LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if pnic.level == passedLogLevel && passedLogLevel >= GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if pnic.next != nil {
		chainErr := pnic.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
