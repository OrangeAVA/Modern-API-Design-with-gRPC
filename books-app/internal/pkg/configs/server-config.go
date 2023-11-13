package configs

type ServerConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	LogLevel    string `mapstructure:"logLevel"`

	ReviewServerAddress   string `mapstructure:"reviewServerAddress"`
	BookServerAddress     string `mapstructure:"bookServerAddress"`
	BookInfoServerAddress string `mapstructure:"bookInfoServerAddress"`
}
