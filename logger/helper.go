package logger

import (
	"os"
)

// Helper ...
type Helper struct {
	Logger
	fields map[string]interface{}
}

// NewHelper ...
func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

// Info ...
func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(InfoLevel, args...)
}

// Infof ...
func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(InfoLevel, template, args...)
}

// Trace ...
func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(TraceLevel, args...)
}

// Tracef ...
func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(TraceLevel, template, args...)
}

// Debug ...
func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(DebugLevel, args...)
}

// Debugf ...
func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(DebugLevel, template, args...)
}

// Warn ...
func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(WarnLevel, args...)
}

// Warnf ...
func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(WarnLevel, template, args...)
}

// Error ...
func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(ErrorLevel, args...)
}

// Errorf ...
func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(ErrorLevel, template, args...)
}

// Fatal ...
func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(FatalLevel, args...)
	os.Exit(1)
}

// Fatalf ...
func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(FatalLevel, template, args...)
	os.Exit(1)
}

// WithError ...
func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

// WithFields ...
func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}
