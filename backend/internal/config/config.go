package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database *DatabaseConfig
	Redis *RedisConfig
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}


type RedisConfig struct {
	Host string
	Port int
	User string
	Password string
}

// Criando nova instancia do Config
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	} 

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		port = 4321
	}

	// Load database information 
	c.Database = &DatabaseConfig{
		Host: os.Getenv("DATABASE_HOST"),
		Port: port,
		User: os.Getenv("DATABASE_PORT"),
		Password: os.Getenv("DATABASE_PASS"),
		DBName: os.Getenv("DATABASE_NAME"),
	}



	return nil
}

