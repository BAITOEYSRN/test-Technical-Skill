package dto

import (
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	"github.com/google/uuid"
)

type ResponseListUser struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
}

func ToResponseListUser(users []entity.User) []ResponseListUser {
	var responseListUsers []ResponseListUser
	for _, user := range users {
		responseListUsers = append(responseListUsers, ResponseListUser{
			ID:          user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			DateOfBirth: user.DateOfBirth,
			Age:         user.Age,
			Address:     user.Address,
		})
	}
	return responseListUsers
}
