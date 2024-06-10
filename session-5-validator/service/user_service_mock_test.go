package service_test

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/entity"
	"time"
)

// MockUserRepository adalah mock untuk IUserRepository
type MockUserRepository struct {
	users []entity.User
}

func (m *MockUserRepository) CreateUser(user *entity.User) entity.User {
	user.ID = len(m.users)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	m.users = append(m.users, *user)
	return *user
}

func (m *MockUserRepository) GetUserByID(id int) (entity.User, bool) {
	for _, user := range m.users {
		if user.ID == id {
			return user, true
		}
	}
	return entity.User{}, false
}

func (m *MockUserRepository) UpdateUser(id int, user entity.User) (entity.User, bool) {
	for i, u := range m.users {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			m.users[i] = user
			return user, true
		}
	}
	return entity.User{}, false
}

func (m *MockUserRepository) DeleteUser(id int) bool {
	for i, user := range m.users {
		if user.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MockUserRepository) GetAllUsers() []entity.User {
	return m.users
}
