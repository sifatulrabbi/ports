package services

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ UsersCRUD = &UsersService{}

type UsersCRUD interface {
	crudService[User, UserPayload, UserFilter]
}

type UsersService struct {
	*ServiceWithDB[User, UserPayload, UserFilter]
}

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email     string    `gorm:"type:text" json:"email"`
	Title     string    `gorm:"type:text" json:"title"`
	Name      string    `gorm:"type:text" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserFilter struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type UserPayload struct {
	Email string `json:"email"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (u *UsersService) CreateOne(p UserPayload) (*User, error) {
	user := &User{
		Email: p.Email,
		Name:  p.Name,
		Title: p.Title,
	}
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUsersService(db *gorm.DB) (*UsersService, error) {
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[User, UserPayload, UserFilter](db, "UsersService")
	usersService := &UsersService{ServiceWithDB: service}
	return usersService, nil
}
