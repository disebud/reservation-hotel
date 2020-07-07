package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

//ConnectDB adalah fungsi untuk koneksi ke database
func Connection() (*sql.DB, error) {

	dbUser := ReadGoEnvVar("dbUser")
	dbPass := ReadGoEnvVar("dbPass")
	dbHost := ReadGoEnvVar("dbHost")
	dbPort := ReadGoEnvVar("dbPort")
	dbName := ReadGoEnvVar("dbName")

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, _ := sql.Open("mysql", source)
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected")

	return db, nil
}

//ReadGoEnvVar adalah fungsi untuk load .env file menggunakan godotenv package
func ReadGoEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
