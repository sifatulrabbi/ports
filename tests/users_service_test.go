package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sifatulrabbi/ports/services"
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

func getUsersService(t *testing.T) *services.UsersService {
	db := getTestDB(t)
	if s, err := services.NewUsersService(db); err != nil {
		t.Error(err.Error())
		return nil
	} else {
		return s
	}
}

func TestCreateUser(t *testing.T) {
	s := getUsersService(t)
	p := services.UserPayload{
		Email: "sifatuli.r@gmail.com",
		Name:  "Sifatul Rabbi",
		Title: "Full Stack Developer",
	}
	u, err := s.CreateOne(p)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(u)
}

func TestGetUserById(t *testing.T) {
	t.FailNow()
}

func TestGetManyUsers(t *testing.T) {
	t.FailNow()
}

func TestUpdateOneUser(t *testing.T) {
	t.FailNow()
}

func TestDeleteOneUser(t *testing.T) {
	t.FailNow()
}
