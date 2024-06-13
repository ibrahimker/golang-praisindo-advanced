package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []User
	userRepo := NewUserRepository(mockUserDBInSlice)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// Routes
	SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}

// userRepository adalah implementasi dari IUserRepository yang menggunakan slice untuk menyimpan data pengguna
type userRepository struct {
	db     []User // slice untuk menyimpan data pengguna
	nextID int    // ID berikutnya yang akan digunakan untuk pengguna baru
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db []User) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAllUsers mengembalikan semua pengguna
func (r *userRepository) GetAllUsers() []User {
	return r.db // Kembalikan slice semua pengguna
}

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	GetAllUsers() []User
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	GetAllUsers() []User
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// GetAllUsers mengembalikan semua pengguna yang ada di repository
func (s *userService) GetAllUsers() []User {
	return s.userRepo.GetAllUsers()
}

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine, userHandler IUserHandler) {
	// Mengatur endpoint publik untuk pengguna
	usersPublicEndpoint := r.Group("/users")
	// Rute untuk mendapatkan semua pengguna
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
}
