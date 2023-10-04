package configs

type ClientConfig struct {
	ClientName    string `mapstructure:"clientName"`
	LogLevel      string `mapstructure:"logLevel"`
	ServerAddress string `mapstructure:"serverAddress"`
}
