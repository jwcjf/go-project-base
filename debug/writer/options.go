package writer

// options 可配置参数
type options struct {
	path   string
	suffix string //文件扩展名
	cap    uint
}

func setDefault() options {
	return options{
		path:   "/tmp/base",
		suffix: "log",
	}
}

// Option ...
type Option func(*options)

// WithPathOption ...
func WithPathOption(s string) Option {
	return func(o *options) {
		o.path = s
	}
}

// WithSuffixOption ...
func WithSuffixOption(s string) Option {
	return func(o *options) {
		o.suffix = s
	}
}

// WithCapOption ...
func WithCapOption(n uint) Option {
	return func(o *options) {
		o.cap = n
	}
}
