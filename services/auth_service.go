package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type AuthTokens struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthService struct {
	name         string
	usersService *UsersService
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthServiceWithDb(db *gorm.DB) (*AuthService, error) {
	usersService, err := NewUsersService(db)
	if err != nil {
		return nil, err
	}
	return &AuthService{name: "AuthService", usersService: usersService}, nil
}

func NewAuthService(usersService *UsersService) *AuthService {
	return &AuthService{name: "AuthService", usersService: usersService}
}

func (s *AuthService) Login(p LoginPayload) (*User, *AuthTokens, error) {
	if len(p.Password) < 8 {
		return nil, nil, errors.New("passwords should 8 characters long")
	}
	if arr := strings.Split(p.Email, "@"); len(arr) < 2 || len(arr) > 2 {
		return nil, nil, errors.New("please provide a valid email address")
	}

	user, err := s.usersService.GetOne(UserFilter{Email: p.Email})
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, errors.New("user not found")
	}
	if user.Password != p.Password {
		return nil, nil, errors.New("wrong password")
	}
	user.Password = ""
	authTokens := AuthTokens{}
	return user, &authTokens, nil
}
