package repository

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-sample-separate-layer/step-4/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-sample-separate-layer/step-4/service"
)

// userRepository adalah implementasi dari IUserRepository yang menggunakan slice untuk menyimpan data pengguna
type userRepository struct {
	db     []entity.User // slice untuk menyimpan data pengguna
	nextID int           // ID berikutnya yang akan digunakan untuk pengguna baru
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db []entity.User) service.IUserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAllUsers mengembalikan semua pengguna
func (r *userRepository) GetAllUsers() []entity.User {
	return r.db // Kembalikan slice semua pengguna
}
