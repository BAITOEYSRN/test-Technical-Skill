package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DateOfBirth string
	Age         int
	Address     string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
