package settings

type Config struct {
	Server   ServerConfig     `mapstructure:"server"`
	Database DatabaseSettings `mapstructure:"database"`
	Logger   LoggerConfig     `mapstructure:"logger"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LoggerConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	FileName   string `mapstructure:"file_name"`
	Path       string `mapstructure:"path"`
	DevMode    bool   `mapstructure:"dev_mode"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type DatabaseSettings struct {
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	Host           string `mapstructure:"host"`
	DbName         string `mapstructure:"dbName"`
	Port           string `mapstructure:"port"`
	SslMode        string `mapstructure:"sslMode"`
	TimeZone       string `mapstructure:"timeZone"`
	MaxIdleTime    int    `mapstructure:"maxIdleTime"`
	MaxIdleConns   int    `mapstructure:"maxIdleConns"`
	MaxOpenConns   int    `mapstructure:"maxOpenConns"`
	ConMaxLifetime int    `mapstructure:"conMaxLifetime"`
}
