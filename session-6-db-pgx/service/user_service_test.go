package service_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/service"
	mock_service "github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/test/mock/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	user := &entity.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(*user, nil)

		createdUser, err := userService.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, *user, createdUser)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(entity.User{}, errors.New("failed to create user"))

		createdUser, err := userService.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
		assert.Equal(t, entity.User{}, createdUser)
	})
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	userID := 1
	user := entity.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(user, nil)

		foundUser, err := userService.GetUserByID(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, user, foundUser)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(entity.User{}, errors.New("user not found"))

		foundUser, err := userService.GetUserByID(ctx, userID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
		assert.Equal(t, entity.User{}, foundUser)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	userID := 1
	user := entity.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().UpdateUser(ctx, userID, user).Return(user, nil)

		updatedUser, err := userService.UpdateUser(ctx, userID, user)
		assert.NoError(t, err)
		assert.Equal(t, user, updatedUser)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().UpdateUser(ctx, userID, user).Return(entity.User{}, errors.New("user not found"))

		updatedUser, err := userService.UpdateUser(ctx, userID, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
		assert.Equal(t, entity.User{}, updatedUser)
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	userID := 1

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUser(ctx, userID).Return(nil)

		err := userService.DeleteUser(ctx, userID)
		assert.NoError(t, err)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUser(ctx, userID).Return(errors.New("user not found"))

		err := userService.DeleteUser(ctx, userID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	users := []entity.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().GetAllUsers(ctx).Return(users, nil)

		retrievedUsers, err := userService.GetAllUsers(ctx)
		assert.NoError(t, err)
		assert.Equal(t, users, retrievedUsers)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().GetAllUsers(ctx).Return([]entity.User{}, errors.New("no users found"))

		retrievedUsers, err := userService.GetAllUsers(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no users found")
		assert.Empty(t, retrievedUsers)
	})
}
