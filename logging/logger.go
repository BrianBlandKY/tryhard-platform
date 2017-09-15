package logging

import (
	"log"
)

type Logger interface {
	Printf(format string, params ...interface{})
	Println(params ...interface{})
	Fatalf(format string, params ...interface{})
	Verboseln(condition bool, params ...interface{})
	Verbosef(condition bool, format string, params ...interface{})
}

func DefaultLogger() Logger {
	return logger{}
}

type logger struct {
}

func (l logger) Printf(format string, params ...interface{}) {
	log.Printf(format, params...)
}

func (l logger) Fatalf(format string, params ...interface{}) {
	log.Fatalf(format, params...)
}

func (l logger) Println(params ...interface{}) {
	log.Println(params...)
}

func (l logger) Verboseln(condition bool, params ...interface{}) {
	if condition {
		l.Println(params...)
	}
}

func (l logger) Verbosef(condition bool, format string, params ...interface{}) {
	if condition {
		l.Printf(format, params...)
	}
}
