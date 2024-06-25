package postgres_gorm

import (
	"context"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/service"
)

type userRepository struct {
	db GormDBIface
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db GormDBIface) service.IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
