package config

import (
	"encoding/json"
	"fmt"
	"os"
	"yukicoding/voteHub/pkg/logger"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service    ServiceConfig    `yaml:"service"`
	Redis      RedisConfig      `yaml:"redis"`
	PostgreSQL PostgreSQLConfig `yaml:"postgresql"`
	Log        LogConfig        `yaml:"log"`
}

type ServiceConfig struct {
	AppMode  string `yaml:"AppMode"`
	HttpPort string `yaml:"HttpPort"`
}

type RedisConfig struct {
	RedisDb     string `yaml:"RedisDb"`
	RedisAddr   string `yaml:"RedisAddr"`
	RedisPw     string `yaml:"RedisPw"`
	RedisDbName string `yaml:"RedisDbName"`
}

type PostgreSQLConfig struct {
	Db              string `yaml:"Db"`
	DbHost          string `yaml:"DbHost"`
	DbPort          int    `yaml:"DbPort"`
	DbUser          string `yaml:"DbUser"`
	DbPassWord      string `yaml:"DbPassWord"`
	DbName          string `yaml:"DbName"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	ConnMaxLifetime string `yaml:"ConnMaxLifetime"`
	SSLMode         string `yaml:"SSLMode"`
	TimeZone        string `yaml:"TimeZone"`
}

// log type
type LogConfig struct {
	LogPath  string `yaml:"LogPath"`
	LogLevel string `yaml:"LogLevel"`
}

func (c *Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("Error marshaling config: %v", err)
	}
	return string(b)
}

func LoadConfig(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(buf))
	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return config, nil
}

// GetPostgreSQLDSN returns the PostgreSQL connection string
func (c *Config) GetPostgreSQLDSN() string {
	logger.Warn(c.PostgreSQL.DbPassWord)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.PostgreSQL.DbHost,
		c.PostgreSQL.DbPort,
		c.PostgreSQL.DbUser,
		c.PostgreSQL.DbPassWord,
		c.PostgreSQL.DbName,
		c.PostgreSQL.SSLMode)

}

// GetRedisAddr returns the Redis address
func (c *Config) GetRedisAddr() string {
	return c.Redis.RedisAddr
}
