package zap

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"git.code.oa.com/ad/go-project-base/logger"
)

// Options ...
type Options struct {
	logger.Options
}

type callerSkipKey struct{}

// WithCallerSkip ...
func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

// WithNamespace ...
func WithNamespace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}

type writerKey struct{}

// WithOutput ...
func WithOutput(out io.Writer) logger.Option {
	return logger.SetOption(writerKey{}, out)
}
