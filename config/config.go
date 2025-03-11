package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// Config is App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   Logger
}

// ServerConfig is server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Debug             bool
}

// PostgresConfig is Postgresql config struct
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// Logger is logger config struct
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// LoadConfig loads configuration from .env file into viper
func LoadConfig() (*viper.Viper, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading .env file")
	}

	v := viper.New()
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig parses environment variables into the Config struct
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	c.Server.AppVersion = v.GetString("SERVER_APPVERSION")
	c.Server.Port = v.GetString("SERVER_PORT")
	c.Server.Mode = v.GetString("SERVER_MODE")
	c.Server.ReadTimeout = time.Duration(v.GetInt("SERVER_READTIMEOUT")) * time.Second
	c.Server.WriteTimeout = time.Duration(v.GetInt("SERVER_WRITETIMEOUT")) * time.Second
	c.Server.CtxDefaultTimeout = time.Duration(v.GetInt("SERVER_CTXDEFAULTTIMEOUT")) * time.Second
	c.Server.Debug = v.GetBool("SERVER_DEBUG")

	c.Postgres.PostgresqlHost = v.GetString("POSTGRES_POSTGRESQLHOST")
	c.Postgres.PostgresqlPort = v.GetString("POSTGRES_POSTGRESQLPORT")
	c.Postgres.PostgresqlUser = v.GetString("POSTGRES_POSTGRESQLUSER")
	c.Postgres.PostgresqlPassword = v.GetString("POSTGRES_POSTGRESQLPASSWORD")
	c.Postgres.PostgresqlDbname = v.GetString("POSTGRES_POSTGRESQLDBNAME")
	c.Postgres.PostgresqlSSLMode = v.GetBool("POSTGRES_POSTGRESQLSSLMODE")
	c.Postgres.PgDriver = v.GetString("POSTGRES_PGDRIVER")

	c.Logger.Development = v.GetBool("LOGGER_DEVELOPMENT")
	c.Logger.DisableCaller = v.GetBool("LOGGER_DISABLECALLER")
	c.Logger.DisableStacktrace = v.GetBool("LOGGER_DISABLESTACKTRACE")
	c.Logger.Encoding = v.GetString("LOGGER_ENCODING")
	c.Logger.Level = v.GetString("LOGGER_LEVEL")

	return &c, nil
}
