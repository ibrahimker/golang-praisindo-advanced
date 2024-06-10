package slice

import (
	"context"
	"errors"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/service"
	"time"
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

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.db = append(r.db, *user)
	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	for _, user := range r.db {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	for i, u := range r.db {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			r.db[i] = user
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	for i, user := range r.db {
		if user.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return r.db, nil
}
