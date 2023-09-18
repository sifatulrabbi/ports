package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sifatulrabbi/ports/services"
)

const userId = "da525320-43e0-478a-a3c2-b5424a6f8fa5"

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

func TestGetUserByEmail(t *testing.T) {
	s := getUsersService(t)
	if user, err := s.GetByEmail(services.UserFilter{Email: "sifatuli.r@gmail.com"}); err != nil {
		t.Error(err)
	} else {
		t.Log(user.String())
	}
}

// when the requested user don't have an account on the database, the backend should reply with a message that clearly states that the user is not found in the db. Later the frontend will send the user to the onboarding page for creating a new user profile for the user.
func TestWhenUserIsNotFound(t *testing.T) {
	s := getUsersService(t)
	if user, err := s.GetByEmail(services.UserFilter{Email: "no-user@example.com"}); err == nil {
		t.Error("backend is not responding with an error even though the user is not found", user)
	} else {
		msg := err.Error()
		if msg != "user not found" {
			t.Error("vague error message", msg)
		}
	}
}
