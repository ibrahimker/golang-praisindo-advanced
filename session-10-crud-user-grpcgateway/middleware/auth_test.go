package middleware_test

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/middleware"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware(t *testing.T) {
	// Set the gin mode to test mode
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		username       string
		password       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid credentials",
			username:       "user",
			password:       "pass",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "Invalid credentials",
			username:       "user",
			password:       "wrongpass",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid authorization token"}`,
		},
		{
			name:           "No credentials",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Authorization basic token required"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Create a new gin router with the AuthMiddleware applied
			router := gin.New()
			router.Use(middleware.AuthMiddleware())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "OK")
			})

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.username != "" || tt.password != "" {
				req.SetBasicAuth(tt.username, tt.password)
			}

			// Create a response recorder to capture the response
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Assert the response status code
			require.Equal(t, tt.expectedStatus, w.Code)

			// Assert the response body based on the type of expected response
			if tt.expectedStatus == http.StatusOK {
				require.Equal(t, tt.expectedBody, w.Body.String())
			} else {
				require.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
