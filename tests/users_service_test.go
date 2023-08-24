package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sifatulrabbi/ports/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const userId = "da525320-43e0-478a-a3c2-b5424a6f8fa5"

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
	t.Log(u.String())
}

func TestGetUserById(t *testing.T) {
	id, err := uuid.Parse(userId)
	if err != nil {
		t.Error(err)
		return
	}
	s := getUsersService(t)
	user, err := s.GetOne(services.UserFilter{ID: id})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user.String())
}

func TestGetManyUsers(t *testing.T) {
	s := getUsersService(t)
	users, err := s.GetMany(services.UserFilter{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("users found %v", len(*users))
}

func TestUpdateOneUser(t *testing.T) {
	id, err := uuid.Parse(userId)
	if err != nil {
		t.Error(err)
		return
	}
	s := getUsersService(t)
	filter := services.UserFilter{ID: id}
	payload := services.UserPayload{
		Title: "Full Stack Developer",
	}
	user, err := s.UpdateOne(filter, payload)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user.String())
}

func TestDeleteOneUser(t *testing.T) {
	id, err := uuid.Parse(userId)
	if err != nil {
		t.Error(err)
		return
	}
	s := getUsersService(t)
	filter := services.UserFilter{ID: id}
	if err = s.DeleteOne(filter); err != nil {
		t.Error(err)
	}
}
