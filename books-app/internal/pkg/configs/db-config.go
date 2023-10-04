package configs

type DatabaseConfig struct {
	Dbname        string         `mapstructure:"name"`
	Schema        string         `mapstructure:"schema"`
	Username      string         `mapstructure:"user"`
	Password      string         `mapstructure:"password"`
	Host          string         `mapstructure:"host"`
	Port          int            `mapstructure:"port"`
	LogMode       bool           `mapstructure:"logMode"`
	SslMode       string         `mapstructure:"sslMode"`
	Connection    ConnectionPool `mapstructure:"connectionPool"`
	MigrationPath string         `mapstructure:"migrationPath"`
}

type ConnectionPool struct {
	MaxOpenConnections int `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int `mapstructure:"maxIdleConnections"`
	MaxIdleTime        int `mapstructure:"maxIdleTime"`
	MaxLifeTime        int `mapstructure:"maxLifeTime"`
	TimeOut            int `mapstructure:"timeout"`
}
