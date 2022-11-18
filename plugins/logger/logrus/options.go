package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/jwcjf/go-project-base/logger"
)

// Options ...
type Options struct {
	logger.Options
	Formatter logrus.Formatter
	Hooks     logrus.LevelHooks
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type formatterKey struct{}

// WithTextTextFormatter ...
func WithTextTextFormatter(formatter *logrus.TextFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

// WithJSONFormatter ...
func WithJSONFormatter(formatter *logrus.JSONFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

type hooksKey struct{}

// WithLevelHooks ...
func WithLevelHooks(hooks logrus.LevelHooks) logger.Option {
	return logger.SetOption(hooksKey{}, hooks)
}

type reportCallerKey struct{}

// warning to use this option. because logrus doest not open CallerDepth option
// this will only print this package
func ReportCaller() logger.Option {
	return logger.SetOption(reportCallerKey{}, true)
}

type exitKey struct{}

// WithExitFunc ...
func WithExitFunc(exit func(int)) logger.Option {
	return logger.SetOption(exitKey{}, exit)
}

type logrusLoggerKey struct{}

// WithLogger ...
func WithLogger(l logrus.StdLogger) logger.Option {
	return logger.SetOption(logrusLoggerKey{}, l)
}
