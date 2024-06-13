package postgres_gorm_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/golang-praisindo-advanced/session-7-db-pg-gorm/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-7-db-pg-gorm/repository/postgres_gorm"
	mock_postgres_gorm "github.com/ibrahimker/golang-praisindo-advanced/session-7-db-pg-gorm/test/mock/repository"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_postgres_gorm.NewMockGormDBIface(ctrl)
	repo := postgres_gorm.NewUserRepository(mockDB)

	t.Run("Positive", func(t *testing.T) {
		user := &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "secret",
		}

		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)

		mockDB.EXPECT().
			Create(user).
			Return(&gorm.DB{Error: nil})

		createdUser, err := repo.CreateUser(context.Background(), user)
		require.NoError(t, err)
		require.NotNil(t, createdUser.ID)
		require.Equal(t, "John Doe", createdUser.Name)
	})

	t.Run("Negative", func(t *testing.T) {
		user := &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "secret",
		}

		mockDB.EXPECT().
			WithContext(gomock.Any()).
			Return(mockDB)

		mockDB.EXPECT().
			Create(user).
			Return(&gorm.DB{Error: errors.New("database error")})

		_, err := repo.CreateUser(context.Background(), user)
		require.Error(t, err)
		require.EqualError(t, err, "database error")
	})
}

func TestUserRepository_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_postgres_gorm.NewMockGormDBIface(ctrl)
	repo := postgres_gorm.NewUserRepository(mockDB)

	t.Run("Positive", func(t *testing.T) {
		// Prepare expected user
		expectedUser := entity.User{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "secret",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// Expect a call to First method on mock with expectedUser.ID as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any(), expectedUser.ID).SetArg(0, expectedUser).Return(&gorm.DB{Error: nil}).Times(1)

		user, err := repo.GetUserByID(context.Background(), expectedUser.ID)
		require.NoError(t, err)
		require.Equal(t, expectedUser.ID, user.ID)
		require.Equal(t, expectedUser.Name, user.Name)
		// Ensure other fields are correctly set
	})

	t.Run("Negative", func(t *testing.T) {
		nonExistentUserID := 999
		// Expect a call to First method on mock and return an error (record not found)
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any(), nonExistentUserID).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)

		_, err := repo.GetUserByID(context.Background(), nonExistentUserID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "record not found")
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_postgres_gorm.NewMockGormDBIface(ctrl)
	repo := postgres_gorm.NewUserRepository(mockDB)

	t.Run("Positive", func(t *testing.T) {
		userID := 1
		user := entity.User{
			ID:        userID,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "newpassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// Expect a call to First method on mock with userID as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any(), userID).SetArg(0, user).Return(&gorm.DB{Error: nil}).Times(1)
		// Expect a call to Save method on mock with updatedUser as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Save(&user).Return(&gorm.DB{}).Times(1)

		updatedUser, err := repo.UpdateUser(context.Background(), userID, user)
		require.NoError(t, err)
		require.Equal(t, user.ID, updatedUser.ID)
		require.Equal(t, user.Name, updatedUser.Name)
		// Ensure other fields are correctly set
	})

	t.Run("Success on first but error on save", func(t *testing.T) {
		userID := 1
		user := entity.User{
			ID:        userID,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "newpassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// Expect a call to First method on mock with userID as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any(), userID).SetArg(0, user).Return(&gorm.DB{Error: nil}).Times(1)
		// Expect a call to Save method on mock with updatedUser as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Save(&user).Return(&gorm.DB{Error: errors.New("db error")}).Times(1)

		updatedUser, err := repo.UpdateUser(context.Background(), userID, user)
		require.Error(t, err)
		require.Empty(t, updatedUser)
		// Ensure other fields are correctly set
	})

	t.Run("Negative", func(t *testing.T) {
		userID := 1
		user := entity.User{
			ID:        userID,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "newpassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// Expect a call to First method on mock with userID as argument and return error
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any(), userID).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)

		_, err := repo.UpdateUser(context.Background(), userID, user)
		require.Error(t, err)
		require.Contains(t, err.Error(), "record not found")
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_postgres_gorm.NewMockGormDBIface(ctrl)
	repo := postgres_gorm.NewUserRepository(mockDB)

	t.Run("Positive", func(t *testing.T) {
		userID := 1
		// Expect a call to Delete method on mock with userID as argument
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Delete(&entity.User{}, userID).Return(&gorm.DB{}).Times(1)

		err := repo.DeleteUser(context.Background(), userID)
		require.NoError(t, err)
	})

	t.Run("Negative", func(t *testing.T) {
		userID := 1
		// Expect a call to Delete method on mock with userID as argument and return error
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Delete(&entity.User{}, userID).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)

		err := repo.DeleteUser(context.Background(), userID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "record not found")
	})
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_postgres_gorm.NewMockGormDBIface(ctrl)
	repo := postgres_gorm.NewUserRepository(mockDB)

	t.Run("Positive", func(t *testing.T) {
		// Prepare expected users
		expectedUsers := []entity.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}
		// Expect a call to Find method on mock and return expectedUsers
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Find(gomock.Any(), gomock.Any()).SetArg(0, expectedUsers).Return(&gorm.DB{Error: nil}).Times(1)

		users, err := repo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.NotNil(t, users)
		require.Len(t, users, len(expectedUsers))
		// Ensure each user matches the expected user
		for i, u := range users {
			require.Equal(t, expectedUsers[i].ID, u.ID)
			require.Equal(t, expectedUsers[i].Name, u.Name)
			require.Equal(t, expectedUsers[i].Email, u.Email)
			// Ensure other fields are correctly set
		}
	})

	t.Run("Negative", func(t *testing.T) {
		// Expect a call to Find method on mock and return error
		mockDB.EXPECT().WithContext(gomock.Any()).Return(mockDB)
		mockDB.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)

		_, err := repo.GetAllUsers(context.Background())
		require.Error(t, err)
		require.Contains(t, err.Error(), "record not found")
	})
}
