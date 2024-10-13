package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null" validate:"required,min=3,max=255"` // Name is required with a length between 3 and 255 characters
	Email    string `gorm:"unique;not null" validate:"required,email"`           // Email must be valid
	Password string `gorm:"not null" validate:"required,min=8"`                  // Password must be at least 8 characters long
}
