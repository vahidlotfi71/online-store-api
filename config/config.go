package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DB   DBConfig
	JWT  JWTConfig
	SMS  SMSConfig
}

type DBConfig struct {
	Host, Port, User, Password, Name, Charset string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type SMSConfig struct {
	APIKey, Sender string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system variables")
	}
	return &Config{
		Port: getEnv("PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "shop_go"),
			Charset:  getEnv("DB_CHARSET", "utf8mb4"),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "secret"),
			ExpireHours: 24,
		},
		SMS: SMSConfig{
			APIKey: getEnv("SMS_API_KEY", ""),
			Sender: getEnv("SMS_SENDER", ""),
		},
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
