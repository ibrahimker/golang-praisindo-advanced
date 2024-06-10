package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`                                // ID pengguna
	Name string `json:"name" binding:"required"`                   // Nama pengguna (wajib diisi)
	Email     string    `json:"email" binding:"required,email"`    // Email pengguna (wajib diisi, harus berformat email)
	Password  string    `json:"password" binding:"required,min=6"` // Kata sandi pengguna (wajib diisi, minimal 6 karakter)
	CreatedAt time.Time `json:"created_at"`                        // Waktu pembuatan pengguna
	UpdatedAt time.Time `json:"updated_at"`                        // Waktu pembaruan terakhir pengguna
}
