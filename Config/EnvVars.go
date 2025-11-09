package Config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_CHARSET  string
	JWT_SECRET  string
	SMS_API_KEY string
	SMS_SENDER  string
)

func Getenv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_CHARSET = os.Getenv("DB_CHARSET")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	SMS_API_KEY = os.Getenv("SMS_API_KEY")
	SMS_SENDER = os.Getenv("SMS_SENDER")

	return nil

}
