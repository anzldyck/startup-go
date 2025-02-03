package user

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// ID              int
	// Name            string
	// Email           string
	// EmailVerifiedAt string
	// Password        string
	// RememberToken   string
	// CreatedAt       time.Time
	// UpdatedAt       time.Time
}
