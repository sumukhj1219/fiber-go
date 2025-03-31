package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error in connecting to db")
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error in pinging db")
	}

	fmt.Println("Database connected sucessfully")
}
