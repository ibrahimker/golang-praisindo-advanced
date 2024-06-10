package slice

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/service"
	"time"
)

type userRepository struct {
	db     []entity.User
	nextID int
}

func NewUserRepository(db []entity.User) service.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entity.User) entity.User {
	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.db = append(r.db, *user)
	return *user
}

func (r *userRepository) GetUserByID(id int) (entity.User, bool) {
	for _, user := range r.db {
		if user.ID == id {
			return user, true
		}
	}
	return entity.User{}, false
}

func (r *userRepository) UpdateUser(id int, user entity.User) (entity.User, bool) {
	for i, u := range r.db {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			r.db[i] = user
			return user, true
		}
	}
	return entity.User{}, false
}

func (r *userRepository) DeleteUser(id int) bool {
	for i, user := range r.db {
		if user.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return true
		}
	}
	return false
}

func (r *userRepository) GetAllUsers() []entity.User {
	return r.db
}
