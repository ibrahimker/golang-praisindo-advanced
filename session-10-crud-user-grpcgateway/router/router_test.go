package router_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/router"
	mock_handler "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/test/mock/handler"
)

func TestSetupRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserHandler := mock_handler.NewMockIUserHandler(ctrl)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Mock middleware to always allow requests for public endpoints
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	router.SetupRouter(r, mockUserHandler)

	// Helper function to create Basic Auth header
	createBasicAuthHeader := func(user, password string) string {
		auth := user + ":" + password
		return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	}

	// Public endpoints
	t.Run("GetUser", func(t *testing.T) {
		mockUserHandler.EXPECT().GetUser(gomock.Any())

		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		mockUserHandler.EXPECT().GetAllUsers(gomock.Any())

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GetAllUsersWithSlash", func(t *testing.T) {
		mockUserHandler.EXPECT().GetAllUsers(gomock.Any())

		req := httptest.NewRequest(http.MethodGet, "/users/", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	// Private endpoints (Unauthorized)
	t.Run("UnauthorizedCreateUser", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("UnauthorizedCreateUserWithSlash", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("UnauthorizedUpdateUser", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("UnauthorizedDeleteUser", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	// Private endpoints (Authorized)
	t.Run("CreateUser", func(t *testing.T) {
		mockUserHandler.EXPECT().CreateUser(gomock.Any())

		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		req.Header.Set("Authorization", createBasicAuthHeader("user", "pass"))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("CreateUserWithSlash", func(t *testing.T) {
		mockUserHandler.EXPECT().CreateUser(gomock.Any())

		req := httptest.NewRequest(http.MethodPost, "/users/", nil)
		req.Header.Set("Authorization", createBasicAuthHeader("user", "pass"))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		mockUserHandler.EXPECT().UpdateUser(gomock.Any())

		req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
		req.Header.Set("Authorization", createBasicAuthHeader("user", "pass"))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		mockUserHandler.EXPECT().DeleteUser(gomock.Any())

		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		req.Header.Set("Authorization", createBasicAuthHeader("user", "pass"))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})
}
