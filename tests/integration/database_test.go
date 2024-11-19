package integration

import (
	"testing"
	"weblazyteam-api/database"
)

func TestDatabaseConnection(t *testing.T) {
	database.Connect()
	if database.DB == nil {
		t.Fatal("Подключение к базе данных не удалось")
	}
}
