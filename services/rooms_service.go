package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ RoomsCRUD = &RoomsService{}

func NewRoomsService(db *gorm.DB) (*RoomsService, error) {
	if err := db.AutoMigrate(&Room{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[Room, RoomPayload, RoomFilter](db, "RoomsService")
	roomsService := &RoomsService{ServiceWithDB: service}
	return roomsService, nil
}

type RoomsCRUD interface {
	crudService[Room, RoomPayload, RoomFilter]
	AddParticipants(filter RoomFilter, userIds []string) (*Room, error)
	RemoveParticipants(filter RoomFilter, userIds []string) (*Room, error)
}

type RoomsService struct {
	*ServiceWithDB[Room, RoomPayload, RoomFilter]
}

type Room struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;not null" json:"id"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null" json:"organization_id"`
	ParticipantIDs UUIDArray `gorm:"type:uuid[];not null" json:"participant_ids"`
	Name           string    `gorm:"type:text;not null" json:"name"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (r *Room) String() string {
	return fmt.Sprintf(
		"Room <%s, %s, %s, %s, %s, %s>\n",
		r.ID, r.Name, r.OrganizationID, r.ParticipantIDs, r.CreatedAt, r.UpdatedAt,
	)
}

type RoomFilter struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
}

type RoomPayload struct {
	OrganizationID string   `json:"organization_id"`
	ParticipantIDs []string `json:"participant_ids"`
	Name           string   `json:"name"`
}

func (s *RoomsService) GetOne(f RoomFilter) (*Room, error) {
	id, err := uuid.Parse(f.ID)
	if err != nil {
		return nil, err
	}
	room := Room{ID: id}
	if err = s.db.First(&room, id).Error; err != nil {
		return nil, err
	}
	if room.Name == "" {
		return nil, fmt.Errorf("no room found with id: %s", id.String())
	}
	return &room, nil
}

func (s *RoomsService) GetMany(filter RoomFilter) (*[]Room, error) {
	rooms := []Room{}
	if filter.OrganizationID == "" {
		return nil, errors.New("please provide organization's id")
	}
	if err := s.db.Where("organization_id = ?", filter.OrganizationID).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (s *RoomsService) CreateOne(p RoomPayload) (*Room, error) {
	orgId, err := uuid.Parse(p.OrganizationID)
	if err != nil {
		return nil, err
	}
	participantIds := UUIDArray{}
	err = participantIds.ParseStringArr(&p.ParticipantIDs)
	if err != nil {
		return nil, err
	}
	room := &Room{
		ID:             uuid.New(),
		Name:           p.Name,
		OrganizationID: orgId,
		ParticipantIDs: participantIds,
	}
	if err := s.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) DeleteOne(filter RoomFilter) error {
	roomId, err := uuid.Parse(filter.ID)
	if err != nil {
		return err
	}
	err = s.db.Delete(&Room{ID: roomId}).Error
	return err
}

func (s *RoomsService) AddParticipants(filter RoomFilter, userIds []string) (*Room, error) {
	roomId, err := uuid.Parse(filter.ID)
	if err != nil {
		return nil, err
	}
	room := Room{ID: roomId}
	if err = s.db.First(&Room{}).Error; err != nil {
		return nil, err
	}
	newIdsArr := room.ParticipantIDs.GetStringArr()
	for _, id := range userIds {
		for _, v := range newIdsArr {
			if v != id {
				newIdsArr = append(newIdsArr, id)
			}
		}
	}
	newUUIDArr := UUIDArray{}
	newUUIDArr.ParseStringArr(&newIdsArr)
	room.ParticipantIDs = newUUIDArr
	if err = s.db.Save(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *RoomsService) RemoveParticipants(filter RoomFilter, userIds []string) (*Room, error) {
	roomId, err := uuid.Parse(filter.ID)
	if err != nil {
		return nil, err
	}
	room := Room{ID: roomId}
	if err = s.db.First(&room).Error; err != nil {
		return nil, err
	}
	newParticipantIds := []string{}
	for _, userId := range userIds {
		for _, v := range room.ParticipantIDs.GetStringArr() {
			if userId != v {
				newParticipantIds = append(newParticipantIds, userId)
			}
		}
	}
	newUUIDArr := UUIDArray{}
	newUUIDArr.ParseStringArr(&newParticipantIds)
	if err = s.db.Save(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
