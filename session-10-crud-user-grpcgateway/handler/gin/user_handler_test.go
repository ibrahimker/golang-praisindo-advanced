package gin_test

import (
	"errors"
	gin2 "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/handler/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/entity"
	mock_service "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/test/mock/service"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := gin2.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}).Return(entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}, nil)

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"John Doe","email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/users", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusCreated, resp.Code)
		require.JSONEq(t, `{"id":0,"name":"John Doe","email":"john@example.com","password":"password","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, resp.Body.String())
	})

	t.Run("InvalidPayload_MissingName", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/users", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"name is mandatory"}`, resp.Body.String())
	})

	t.Run("InvalidPayload_MissingEmail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"john","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/users", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"email is mandatory"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}).Return(entity.User{}, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"John Doe","email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/users", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := gin2.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(gomock.Any(), 1).Return(entity.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}, nil)

		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/users/:id", userHandler.GetUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.JSONEq(t, `{"id":1,"name":"John Doe","email":"john@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, resp.Body.String())
	})

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/users/:id", userHandler.GetUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(gomock.Any(), 1).Return(entity.User{}, errors.New("user not found"))

		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/users/:id", userHandler.GetUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"User not found"}`, resp.Body.String())
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := gin2.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().UpdateUser(gomock.Any(), 1, entity.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}).Return(entity.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}, nil)

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/users/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.JSONEq(t, `{"id":1,"name":"John Doe","email":"john@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, resp.Body.String())
	})

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users/abc", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/users/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("InvalidPayload", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/users/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"name is mandatory"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().UpdateUser(gomock.Any(), 1, entity.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}).Return(entity.User{}, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/users/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := gin2.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().DeleteUser(gomock.Any(), 1).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/users/:id", userHandler.DeleteUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.JSONEq(t, `{"message":"User deleted"}`, resp.Body.String())
	})

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/abc", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/users/:id", userHandler.DeleteUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().DeleteUser(gomock.Any(), 1).Return(errors.New("some service error"))

		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/users/:id", userHandler.DeleteUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := gin2.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().GetAllUsers(gomock.Any()).Return([]entity.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}, nil)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/users", userHandler.GetAllUsers)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.JSONEq(t, `[{"id":1,"name":"John Doe","email":"john@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},{"id":2,"name":"Jane Doe","email":"jane@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().GetAllUsers(gomock.Any()).Return(nil, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/users", userHandler.GetAllUsers)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}
