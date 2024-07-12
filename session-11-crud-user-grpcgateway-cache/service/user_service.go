package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ibrahimker/golang-praisindo-advanced/session-11-crud-user-grpcgateway-cache/entity"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
	rdb      *redis.Client
}

const redisUserByIDKey = "user:%d"

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository, rdb *redis.Client) IUserService {
	return &userService{userRepo: userRepo, rdb: rdb}
}

// CreateUser membuat pengguna baru
func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}

	fmt.Println(createdUser)
	// simpan di cache
	redisKey := fmt.Sprintf(redisUserByIDKey, createdUser.ID)
	createdUserJSON, err := json.Marshal(createdUser)
	if err != nil {
		log.Println("gagal marshal json")
	}
	if err := s.rdb.Set(ctx, redisKey, createdUserJSON, 60*time.Second).Err(); err != nil {
		log.Println("error when set redis key", redisKey)
	}
	return createdUser, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	// cek data cache, kalau ada, ambil dari redis. kalau tidak ada, ambil dari repo
	var user entity.User
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	val, err := s.rdb.Get(ctx, redisKey).Result()
	if err == nil {
		log.Println("data tersedia di redis")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Println("gagal marshall data di redis, coba ambil ke database")
		}
	}
	if err != nil {
		log.Println("data tidak tersedia di redis, ambil dari database")
		user, err = s.userRepo.GetUserByID(ctx, id)
		if err != nil {
			log.Println("gagal ambil data di database")
			return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
		}
	}

	return user, nil
}

// UpdateUser memperbarui data pengguna
func (s *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	// delete redis key
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	if err := s.rdb.Del(ctx).Err(); err != nil {
		log.Println("gagal delete key redis", redisKey)
	}
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}
