package slice_test

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/repository/slice"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository(t *testing.T) {
	repo := slice.NewUserRepository([]entity.User{})

	t.Run("CreateUser", func(t *testing.T) {
		newUser := entity.User{Name: "Bob", Email: "bob@example.com", Password: "password"}
		createdUser := repo.CreateUser(&newUser)

		require.Equal(t, 0, createdUser.ID)
		require.Equal(t, "Bob", createdUser.Name)
		require.Equal(t, "bob@example.com", createdUser.Email)
		require.NotZero(t, createdUser.CreatedAt)
		require.NotZero(t, createdUser.UpdatedAt)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		user, found := repo.GetUserByID(0)

		require.True(t, found)
		require.Equal(t, "Bob", user.Name)

		_, notFound := repo.GetUserByID(99)
		require.False(t, notFound)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		update := entity.User{Name: "Bob Updated", Email: "bob.updated@example.com", Password: "newpassword"}
		updatedUser, found := repo.UpdateUser(0, update)

		require.True(t, found)
		require.Equal(t, "Bob Updated", updatedUser.Name)
		require.Equal(t, "bob.updated@example.com", updatedUser.Email)

		_, notFound := repo.UpdateUser(99, update)
		require.False(t, notFound)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		deleted := repo.DeleteUser(0)
		require.True(t, deleted)

		notDeleted := repo.DeleteUser(99)
		require.False(t, notDeleted)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		allUsers := repo.GetAllUsers()
		require.Empty(t, allUsers) // Tidak tersisa pengguna setelah penghapusan
	})
}
