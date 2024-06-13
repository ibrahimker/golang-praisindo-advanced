package service

import "github.com/ibrahimker/golang-praisindo-advanced/session-4-sample-separate-layer/step-3/entity"

// userRepository adalah implementasi dari IUserRepository yang menggunakan slice untuk menyimpan data pengguna
type userRepository struct {
	db     []entity.User // slice untuk menyimpan data pengguna
	nextID int           // ID berikutnya yang akan digunakan untuk pengguna baru
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db []entity.User) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAllUsers mengembalikan semua pengguna
func (r *userRepository) GetAllUsers() []entity.User {
	return r.db // Kembalikan slice semua pengguna
}

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
