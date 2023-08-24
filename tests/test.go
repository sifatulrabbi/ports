package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getTestDB(t *testing.T) *gorm.DB {
	godotenv.Load("../.env")
	pgHost := os.Getenv("PGHOST")
	pgDbName := os.Getenv("PGDBNAME")
	pgSSL := os.Getenv("PGSSL")
	if pgHost == "" || pgDbName == "" || pgSSL == "" {
		t.Error("required env vars not found")
		return nil
	}

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=%s", pgHost, pgDbName, pgSSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Errorf("unable to connect to the db\nerror: %v\n", err)
	}
	return db
}
