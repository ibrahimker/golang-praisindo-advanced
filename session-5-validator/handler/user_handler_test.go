package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/handler"
)

func TestGetAllUserHandler(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.GET("/users", userHandler.GetAllUsers)

		req, _ := http.NewRequest("GET", "/users", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var users []entity.User
		err := json.Unmarshal(w.Body.Bytes(), &users)
		require.NoError(t, err)
		require.Equal(t, 2, len(users))
	})
}

func TestCreateUserHandler(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.POST("/users", userHandler.CreateUser)

		user := entity.User{Name: "Test User", Email: "test@example.com", Password: "testpass"}
		jsonUser, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)

		var createdUser entity.User
		err := json.Unmarshal(w.Body.Bytes(), &createdUser)
		require.NoError(t, err)
		require.Equal(t, user.Name, createdUser.Name)
		require.Equal(t, user.Email, createdUser.Email)
	})

	t.Run("Negative Test Case - invalid json", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}

		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.POST("/users", userHandler.CreateUser)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		var customError struct{ Error string }
		_ = json.Unmarshal(w.Body.Bytes(), &customError)
		require.Equal(t, "invalid character 'i' looking for beginning of value", customError.Error)
	})

	t.Run("Negative Test Case - empty email", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.POST("/users", userHandler.CreateUser)

		user := entity.User{Name: "Test User", Email: "", Password: "testpass"}
		jsonUser, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)

		var customError struct{ Error string }
		_ = json.Unmarshal(w.Body.Bytes(), &customError)
		require.Equal(t, "email is mandatory", customError.Error)
	})
}

func TestGetUserHandler(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.GET("/users/:id", userHandler.GetUser)

		req, _ := http.NewRequest("GET", "/users/1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var user entity.User
		err := json.Unmarshal(w.Body.Bytes(), &user)
		require.NoError(t, err)
		require.Equal(t, 1, user.ID)
	})

	t.Run("Negative Test Case - no data", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.GET("/users/:id", userHandler.GetUser)

		// Using invalid user ID
		req, _ := http.NewRequest("GET", "/users/100", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Negative Test Case - cannot convert to integer", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.GET("/users/:id", userHandler.GetUser)

		// Using invalid user ID
		req, _ := http.NewRequest("GET", "/users/abc", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateUserHandler(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.PUT("/users/:id", userHandler.UpdateUser)

		updatedUser := entity.User{ID: 1, Name: "Updated User", Email: "updated@example.com", Password: "updatedpass"}
		jsonUser, _ := json.Marshal(updatedUser)

		req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var returnedUser entity.User
		err := json.Unmarshal(w.Body.Bytes(), &returnedUser)
		require.NoError(t, err)
		require.Equal(t, updatedUser.Name, returnedUser.Name)
		require.Equal(t, updatedUser.Email, returnedUser.Email)
	})

	t.Run("Negative Test Case - invalid id", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.PUT("/users/:id", userHandler.UpdateUser)

		req, _ := http.NewRequest("PUT", "/users/abc", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Negative Test Case - invalid json", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.PUT("/users/:id", userHandler.UpdateUser)

		req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Negative Test Case - name is empty", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.PUT("/users/:id", userHandler.UpdateUser)

		invalidUser := entity.User{ID: 100, Name: "", Email: "updated@example.com", Password: "updatedpass"} // User with ID 100 doesn't exist
		jsonUser, _ := json.Marshal(invalidUser)

		req, _ := http.NewRequest("PUT", "/users/100", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)

		var customError struct{ Error string }
		_ = json.Unmarshal(w.Body.Bytes(), &customError)
		require.Equal(t, "name is mandatory", customError.Error)
	})

	t.Run("Negative Test Case - not found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.PUT("/users/:id", userHandler.UpdateUser)

		invalidUser := entity.User{ID: 100, Name: "Updated User", Email: "updated@example.com", Password: "updatedpass"} // User with ID 100 doesn't exist
		jsonUser, _ := json.Marshal(invalidUser)

		req, _ := http.NewRequest("PUT", "/users/100", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDeleteUserHandler(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.DELETE("/users/:id", userHandler.DeleteUser)

		req, _ := http.NewRequest("DELETE", "/users/1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Negative Test Case - Cannot Convert Int to String", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.DELETE("/users/:id", userHandler.DeleteUser)

		req, _ := http.NewRequest("DELETE", "/users/abc", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Negative Test Case - Not Found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := &MockUserService{}
		userHandler := handler.NewUserHandler(mockUserService)

		r := gin.Default()
		r.DELETE("/users/:id", userHandler.DeleteUser)

		req, _ := http.NewRequest("DELETE", "/users/100", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}
