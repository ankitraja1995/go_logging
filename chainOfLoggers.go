package go_logging

func GetChainOfLoggers() abstractLogger {
	warn := &warn{}
	warn.setLevel(WarnLevel)

	error := &errs{}
	error.setLevel(ErrorLevel)

	panic := &panic{}
	panic.setLevel(PanicLevel)

	fatal := &fatal{}
	fatal.setLevel(FatalLevel)

	debug := &debug{}
	debug.setLevel(DebugLevel)

	info := &info{}
	info.setLevel(InfoLevel)

	trace := &trace{}
	trace.setLevel(Trace_level)

	warn.setNext(error)
	error.setNext(panic)
	panic.setNext(fatal)
	fatal.setNext(debug)
	debug.setNext(info)
	info.setNext(trace)

	return warn
}
