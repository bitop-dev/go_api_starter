package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config represents the application configuration loaded via Viper.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	DB       DatabaseConfig `mapstructure:"db"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Features FeatureFlags   `mapstructure:"features"`
}

// ServerConfig contains HTTP server options.
type ServerConfig struct {
	Port            int    `mapstructure:"port"`
	ReadTimeoutSec  int    `mapstructure:"read_timeout_sec"`
	WriteTimeoutSec int    `mapstructure:"write_timeout_sec"`
	Env             string `mapstructure:"env"`
}

// DatabaseConfig describes the database connection.
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

// LoggingConfig controls slog behavior.
type LoggingConfig struct {
	Level string `mapstructure:"level"`
	JSON  bool   `mapstructure:"json"`
}

// FeatureFlags toggles optional modules.
type FeatureFlags struct {
	AuthJWT    FeatureToggle `mapstructure:"auth_jwt"`
	CacheRedis FeatureToggle `mapstructure:"cache_redis"`
	QueueNATS  FeatureToggle `mapstructure:"queue_nats"`
}

// FeatureToggle represents a boolean flag with optional metadata.
type FeatureToggle struct {
	Enable bool `mapstructure:"enable"`
}

// Load reads configuration from file and environment variables.
func Load(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if path != "" {
		v.AddConfigPath(path)
	}
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.read_timeout_sec", 15)
	v.SetDefault("server.write_timeout_sec", 15)
	v.SetDefault("server.env", "development")

	v.SetDefault("db.driver", "postgres")
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", 5432)
	v.SetDefault("db.user", "app")
	v.SetDefault("db.pass", "password")
	v.SetDefault("db.name", "appdb")
	v.SetDefault("db.sslmode", "disable")

	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.json", true)

	v.SetDefault("features.auth_jwt.enable", false)
	v.SetDefault("features.cache_redis.enable", false)
	v.SetDefault("features.queue_nats.enable", false)
}
