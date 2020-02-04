package main

type abstractLogger interface {
	Execute(*FluentdLogger, LogLevel, string, map[string]string) error
	setNext(abstractLogger)
}
