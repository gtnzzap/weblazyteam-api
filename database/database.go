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
	connString := "sqlserver://petr:petr@host.docker.internal:1433?database=weblazyteam"
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Database connection established")
}
