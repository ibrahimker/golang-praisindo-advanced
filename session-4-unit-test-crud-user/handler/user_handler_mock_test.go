package handler_test

import (
	"errors"

	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/entity"
)

// MockUserService adalah mock untuk service.IUserService
type MockUserService struct{}

var (
	mockUsers      []entity.User
	mockNextUserID int
)

func init() {
	mockUsers = []entity.User{
		{ID: 1, Name: "User 1", Email: "user1@example.com", Password: "pass1"},
		{ID: 2, Name: "User 2", Email: "user2@example.com", Password: "pass2"},
	}
	mockNextUserID = 3
}

func (m *MockUserService) CreateUser(user *entity.User) entity.User {
	user.ID = mockNextUserID
	mockNextUserID++
	mockUsers = append(mockUsers, *user)
	return *user
}

func (m *MockUserService) GetUserByID(id int) (entity.User, error) {
	for _, user := range mockUsers {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (m *MockUserService) UpdateUser(id int, user entity.User) (entity.User, error) {
	for i, u := range mockUsers {
		if u.ID == id {
			user.ID = id
			mockUsers[i] = user
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (m *MockUserService) DeleteUser(id int) error {
	for i, user := range mockUsers {
		if user.ID == id {
			mockUsers = append(mockUsers[:i], mockUsers[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (m *MockUserService) GetAllUsers() []entity.User {
	return mockUsers
}
