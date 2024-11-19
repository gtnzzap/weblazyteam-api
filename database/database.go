package database

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb" // SQL Server драйвер
)

var DB *sql.DB

// Подключение к базе данных
func Connect() {
	var err error
	connString := "sqlserver://localhost:1433?database=weblazyteam&trusted_connection=true"
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Database connection established")
}
