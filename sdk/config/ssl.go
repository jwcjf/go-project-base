package config

// Ssl ...
type Ssl struct {
	KeyStr string
	Pem    string
	Enable bool
	Domain string
}

// SslConfig ...
var SslConfig = new(Ssl)
