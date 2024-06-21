package service

import (
	"context"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/entity"
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
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
