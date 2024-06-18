package database

import (
	"fmt"
	"github.com/jimmyfish/thalopsistematika/configs"
	"github.com/jimmyfish/thalopsistematika/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	config := configs.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DbHost, config.DbUser, config.DbName, config.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&user.User{})

	return db

}

