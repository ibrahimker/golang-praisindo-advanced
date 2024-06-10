// service/user_service.go
package service

import (
	"fmt"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/entity"
)

type UserService interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, error)
	UpdateUser(id int, user entity.User) (entity.User, error)
	DeleteUser(id int) error
	GetAllUsers() []entity.User
}

type UserRepository interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, bool)
	UpdateUser(id int, user entity.User) (entity.User, bool)
	DeleteUser(id int) bool
	GetAllUsers() []entity.User
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *entity.User) entity.User {
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (entity.User, error) {
	user, found := s.userRepo.GetUserByID(id)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id int, user entity.User) (entity.User, error) {
	updatedUser, found := s.userRepo.UpdateUser(id, user)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(id int) error {
	if !s.userRepo.DeleteUser(id) {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
