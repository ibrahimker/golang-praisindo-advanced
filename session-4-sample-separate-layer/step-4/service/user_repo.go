package service

import "github.com/ibrahimker/golang-praisindo-advanced/session-4-sample-separate-layer/step-4/entity"

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	GetAllUsers() []entity.User
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	GetAllUsers() []entity.User
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// GetAllUsers mengembalikan semua pengguna yang ada di repository
func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
