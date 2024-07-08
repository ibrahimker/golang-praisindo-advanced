package slice_test

import (
	"context"
	"testing"
	"time"

	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/repository/slice"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/service"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	repo := slice.NewUserRepository([]entity.User{})
	userService := service.NewUserService(repo)

	newUser := entity.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	createdUser, err := userService.CreateUser(context.Background(), &newUser)
	require.NoError(t, err)
	require.Equal(t, 0, createdUser.ID)
	require.Equal(t, newUser.Name, createdUser.Name)
	require.Equal(t, newUser.Email, createdUser.Email)
	require.Equal(t, newUser.Password, createdUser.Password)
}

func TestGetUserByID(t *testing.T) {
	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	repo := slice.NewUserRepository(users)
	userService := service.NewUserService(repo)

	user, err := userService.GetUserByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, 1, user.ID)
	require.Equal(t, "John Doe", user.Name)
	require.Equal(t, "john@example.com", user.Email)
	require.Equal(t, "password", user.Password)
}

func TestUpdateUser(t *testing.T) {
	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	repo := slice.NewUserRepository(users)
	userService := service.NewUserService(repo)

	userToUpdate := entity.User{
		Name: "Updated Name",
	}
	updatedUser, err := userService.UpdateUser(context.Background(), 1, userToUpdate)
	require.NoError(t, err)
	require.Equal(t, 1, updatedUser.ID)
	require.Equal(t, "Updated Name", updatedUser.Name)

	// Check if the user was actually updated
	updatedUserFromDB, _ := userService.GetUserByID(context.Background(), 1)
	require.Equal(t, "Updated Name", updatedUserFromDB.Name)
}

func TestDeleteUser(t *testing.T) {
	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	repo := slice.NewUserRepository(users)
	userService := service.NewUserService(repo)

	err := userService.DeleteUser(context.Background(), 1)
	require.NoError(t, err)

	// Check if the user was actually deleted
	_, err = userService.GetUserByID(context.Background(), 1)
	require.Error(t, err)
	require.Contains(t, err.Error(), "user not found")
}

func TestGetAllUsers(t *testing.T) {
	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	repo := slice.NewUserRepository(users)
	userService := service.NewUserService(repo)

	allUsers, err := userService.GetAllUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, allUsers, 2)
}
