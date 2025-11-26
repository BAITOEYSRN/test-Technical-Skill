package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName   string     `gorm:"not null"`
	LastName    string     `gorm:"not null"`
	DateOfBirth string     `gorm:"not null"`
	Age         int        `gorm:"not null"`
	Address     string     `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null;default:now()"`
	UpdatedAt   *time.Time `gorm:"default:null"`
}

func (u *User) TableName() string {
	return "users"
}
