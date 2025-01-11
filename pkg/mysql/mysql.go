package mysql

import (
	"fmt"

	"github.com/gofor-little/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(config *gorm.Config) (*gorm.DB, error) {
	dbHost := env.Get("DB_HOST", "")
	dbPort := env.Get("DB_PORT", "")
	dbDatabase := env.Get("DB_DATABASE", "")
	dbUser := env.Get("DB_USER", "")
	dbPassw := env.Get("DB_PASSW", "")

	if config == nil {
		config = &gorm.Config{}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassw, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
