package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ MessagesCRUD = &MessagesService{}

func NewMessagesService(db *gorm.DB) (*MessagesService, error) {
	if err := db.AutoMigrate(&Message{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[Message, MessagePayload, MessageFilter](db, "MessagesService")
	messagesService := &MessagesService{ServiceWithDB: service}
	return messagesService, nil
}

type MessagesCRUD interface {
	crudService[Message, MessagePayload, MessageFilter]
}

type MessagesService struct {
	*ServiceWithDB[Message, MessagePayload, MessageFilter]
}

type Message struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	SenderID  uuid.UUID `gorm:"type:uuid;not null;reference:ID" json:"sender_id"`
	RoomID    uuid.UUID `gorm:"type:uuid;not null;reference:ID" json:"room_id"`
	Richtext  string    `gorm:"text" json:"richtext"`
	Audio     string    `gorm:"type:text" json:"audio"`
	Video     string    `gorm:"type:text" json:"video"`
	Img       string    `gorm:"type:text" json:"img"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (msg *Message) String() string {
	return fmt.Sprintf(
		"Message <%s, %s, %s, %s, %s, %s, %s, %s, %s>",
		msg.ID,
		msg.SenderID,
		msg.RoomID,
		msg.Richtext,
		msg.Audio,
		msg.Video,
		msg.Img,
		msg.CreatedAt,
		msg.UpdatedAt,
	)
}

type MessageFilter struct {
	ID       uuid.UUID `json:"id"`
	SenderID uuid.UUID `gorm:"type:uuid" json:"sender_id"`
	RoomID   uuid.UUID `gorm:"type:uuid" json:"room_id"`
}

type MessagePayload struct {
	SenderID uuid.UUID `gorm:"type:uuid" json:"sender_id"`
	RoomID   uuid.UUID `gorm:"type:uuid" json:"room_id"`
	Richtext string    `gorm:"text" json:"richtext"`
	Audio    string    `gorm:"type:text" json:"audio"`
	Video    string    `gorm:"type:text" json:"video"`
	Img      string    `gorm:"type:text" json:"img"`
}

func (s *MessagesService) GetOne(p MessageFilter) (*Message, error) {
	msg := Message{}
	res := s.db.First(&msg, p.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &msg, nil
}

func (s *MessagesService) GetMany(p MessageFilter) (*[]Message, error) {
	msgs := []Message{}
	res := s.db.Find(&msgs)
	if res.Error != nil {
		return nil, res.Error
	}
	return &msgs, nil
}

func (s *MessagesService) CreateOne(p MessagePayload) (*Message, error) {
	msg := Message{
		ID:       uuid.New(),
		SenderID: p.SenderID,
		RoomID:   p.RoomID,
		Richtext: p.Richtext,
		Audio:    p.Audio,
		Video:    p.Video,
		Img:      p.Img,
	}
	if err := s.db.Create(&msg).Error; err != nil {
		return nil, err
	}
	return &msg, nil
}

func (s *MessagesService) UpdateOne(f MessageFilter, p MessagePayload) (*Message, error) {
	msg := Message{ID: f.ID}
	res := s.db.Model(&msg).Where("id = ?", f.ID).Updates(p)
	if res.Error != nil {
		return nil, res.Error
	}
	if err := res.First(&msg, f.ID).Error; err != nil {
		return nil, err
	}
	return &msg, nil
}

func (s *MessagesService) DeleteOne(f MessageFilter) error {
	msg := Message{ID: f.ID}
	res := s.db.Delete(&msg, f.ID)
	return res.Error
}
