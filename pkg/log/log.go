package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"accounting/pkg/config"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Print(args ...interface{})

	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

type logger struct {
	entry logrus.Entry
}

func NewConfigless() Logger {

	hostname, _ := os.Hostname()

	l := logrus.New()
	l.Formatter = &Formatter{host: hostname}

	return &logger{
		entry: *logrus.NewEntry(l),
	}
}

func NewLogger(config *config.Configuration) Logger {

	hostname, _ := os.Hostname()

	l := logrus.New()
	level := config.Logger.LogLevel
	if level != nil {
		parseLevel, e := logrus.ParseLevel(*level)
		if e == nil {
			l.SetLevel(parseLevel)
		}
	}

	l.Formatter = &Formatter{host: hostname}
	l.SetOutput(os.Stdout)

	return &logger{
		entry: *logrus.NewEntry(l),
	}
}

func (l *logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logger) Print(args ...interface{}) {
	l.entry.Print(args...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.entry.Infof(format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.entry.Warnf(format, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.entry.Errorf(format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if v != nil {
		// RESTY sends this without v interfaces
		l.entry.Debugf(format, v...)
	} else {
		l.entry.Debugf(format)
	}
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.entry.Fatalf(format, v...)
}

func (l *logger) Panicf(format string, v ...interface{}) {
	l.entry.Panicf(format, v...)
}
