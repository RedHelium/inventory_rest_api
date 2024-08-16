package config

import "fmt"

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

func GetConnectionString(config *Config) string {

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
}
