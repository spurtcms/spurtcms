package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	er := godotenv.Load()

	if er != nil {
		log.Fatalf("Error loading .env file")
	}

	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	var (
		db  *gorm.DB
		err error
	)

	DataBaseType := os.Getenv("DATABASE_TYPE")

	switch DataBaseType {

	case "postgres":
		dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")) //Build connection string
		db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	case "mysql":
		dsn := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	default:
		return nil
	}

	if err != nil {
		fmt.Println(err)
	}

	return db
}
