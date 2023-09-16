package tests

import (
	"testing"

	"github.com/sifatulrabbi/ports/services"
)

func getRoomsService(t *testing.T) *services.RoomsService {
	db := getTestDB(t)
	if s, err := services.NewRoomsService(db); err != nil {
		t.Error(err)
		return nil
	} else {
		return s
	}
}

func TestCreateRoom(t *testing.T) {
	s := getRoomsService(t)
	payload := services.RoomPayload{
		Name:           "Test room 1",
		OrganizationID: "0f841993-41fa-4f55-b536-b4f057eaf454",
		ParticipantIDs: []string{"da525320-43e0-478a-a3c2-b5424a6f8fa5"},
	}
	if room, err := s.CreateOne(payload); err != nil {
		t.Error(err)
	} else {
		t.Log(room.String())
	}
}

func TestGetARoom(t *testing.T) {
	s := getRoomsService(t)
	if room, err := s.GetOne(services.RoomFilter{ID: "0d9dc371-7943-4cd2-843b-c03c5ad40e15"}); err != nil {
		t.Error(err)
	} else {
		t.Log(room.String())
	}
}

func TestGetRooms(t *testing.T) {
	s := getRoomsService(t)
	if rooms, err := s.GetMany(services.RoomFilter{OrganizationID: "0f841993-41fa-4f55-b536-b4f057eaf454"}); err != nil {
		t.Error(err)
	} else {
		t.Log("Rooms found:", len(*rooms))
	}
}

func TestAddParticipantsToARoom(t *testing.T) {

}

func TestRemoveParticipantsFromARoom(t *testing.T) {

}

func TestDisbandARoom(t *testing.T) {

}
