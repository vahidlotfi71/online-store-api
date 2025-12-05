package Config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local&sql_mode=STRICT_TRANS_TABLES&interpolateParams=false",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, DB_CHARSET, "utf8mb4_unicode_ci")

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		return fmt.Errorf("config, connect : %s", err.Error())
	}
	DB = db
	return nil
}
