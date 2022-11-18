package config

// Application ...
type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          int64
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
	EnableDP      bool
}

// ApplicationConfig ...
var ApplicationConfig = new(Application)
