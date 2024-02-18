package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	host := os.Getenv("POSTGRES_HOST_INVEST")
	port := os.Getenv("POSTGRES_PORT_INVEST")
	user := os.Getenv("POSTGRES_USER_INVEST")
	dbname := os.Getenv("POSTGRES_DB_INVEST")
	password := os.Getenv("POSTGRES_PASSWORD_INVEST")
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Veritabanına bağlanırken hata oluştu: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("Veritabanına bağlanırken hata oluştu: %v", err))
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Errorf("Veritabanına ping atılırken hata oluştu: %v", err))
	}

	log.Println("Veritabanına başarıyla bağlandı")

	return db
}
