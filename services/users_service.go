package services

import (
	"fmt"
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

func (u *User) String() string {
	return fmt.Sprintf("User<%s, %s, %s, %s, %s, %s>", u.ID, u.Email, u.Name, u.Title, u.CreatedAt, u.UpdatedAt)
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

func (s *UsersService) CreateOne(p UserPayload) (*User, error) {
	user := &User{
		ID:    uuid.New(),
		Email: p.Email,
		Name:  p.Name,
		Title: p.Title,
	}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UsersService) UpdateOne(f UserFilter, p UserPayload) (*User, error) {
	user := User{ID: f.ID}
	res := s.db.Model(&user).Where("id = ?", f.ID).Updates(p)
	if res.Error != nil {
		s.Log("unable to update the user\nerror:%s\n", res.Error.Error())
		return nil, res.Error
	}
	if err := res.First(&user, user.ID); err != nil {
		s.Log("unable to get the updated user\nerror:%s\n", res.Error.Error())
		return nil, res.Error
	}
	return &user, nil
}

func (s *UsersService) DeleteOne(f UserFilter) error {
	user := User{ID: f.ID}
	err := s.db.Delete(&user, f.ID).Error
	return err
}

func (u *UsersService) GetOne(p UserFilter) (*User, error) {
	user := User{}
	if err := u.db.First(&user, p.ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersService) GetMany(p UserFilter) (*[]User, error) {
	users := []User{}
	res := s.db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

func NewUsersService(db *gorm.DB) (*UsersService, error) {
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[User, UserPayload, UserFilter](db, "UsersService")
	usersService := &UsersService{ServiceWithDB: service}
	return usersService, nil
}
