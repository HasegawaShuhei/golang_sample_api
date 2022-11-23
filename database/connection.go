package database

import (
	"fmt"
	"log"
	"os"

	"github.com/HasegawaShuhei/golang_sample_api/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env load failed %v", err)
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	fmt.Println("db connected!!")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatal("CloseDB failed:", err)
	}
}
