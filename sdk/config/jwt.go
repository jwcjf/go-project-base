package config

// Jwt ...
type Jwt struct {
	Secret  string
	Timeout int64
}

// JwtConfig ...
var JwtConfig = new(Jwt)
