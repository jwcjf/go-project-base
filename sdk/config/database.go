package config

// Database ...
type Database struct {
	Driver          string
	Source          string
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
	Registers       []DBResolverConfig
}

// DBResolverConfig ...
type DBResolverConfig struct {
	Sources  []string
	Replicas []string
	Policy   string
	Tables   []string
}

// DatabaseConfig ...
var (
	DatabaseConfig  = new(Database)
	DatabasesConfig = make(map[string]*Database)
)
