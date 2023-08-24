package tests

import (
	"testing"

	"github.com/sifatulrabbi/ports/services"
)

func getMessagesService(t *testing.T) *services.MessagesService {
	db := getTestDB(t)
	if s, err := services.NewMessagesService(db); err != nil {
		t.Error(err.Error())
		return nil
	} else {
		return s
	}
}

func TestGetOneMessage(t *testing.T) {
	s := getMessagesService(t)
	if msg, err := s.GetOne(services.MessageFilter{}); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(msg.String())
	}
}

func TestGetManyMessage(t *testing.T) {
	s := getMessagesService(t)
	if msgs, err := s.GetMany(services.MessageFilter{}); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("messages: %v", len(*msgs))
	}
}

func TestCreateMessage(t *testing.T) {
	s := getMessagesService(t)
	p := services.MessagePayload{
		Richtext: "<p>Hello world! This is a test message from Go tests package.</p>",
	}
	if msg, err := s.CreateOne(p); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(msg.String())
	}
}

func TestUpdateOneMessage(t *testing.T) {
	s := getMessagesService(t)
	f := services.MessageFilter{}
	p := services.MessagePayload{}
	if msg, err := s.UpdateOne(f, p); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(msg.String())
	}
}

func TestDeleteOneMessage(t *testing.T) {
	s := getMessagesService(t)
	f := services.MessageFilter{}
	if err := s.DeleteOne(f); err != nil {
		t.Error(err)
	}
}
