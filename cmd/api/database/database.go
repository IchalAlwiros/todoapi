package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() (err error)  {

	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file:" , er)
	}

	dBHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")


	fmt.Println("DB_HOST:", dBHost)
	fmt.Println("DB_HOST:", dbUser)
	fmt.Println("DB_HOST:", dbName)
	fmt.Println("DB_HOST:", dbPassword)
	fmt.Println("DB_HOST:", dbPort)


	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dBHost, dbUser, dbPassword,dbName, dbPort)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connection to the database:", err)
	}

	return err
}