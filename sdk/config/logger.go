package config

// Logger ...
type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
}

// LoggerConfig ...
var LoggerConfig = new(Logger)
