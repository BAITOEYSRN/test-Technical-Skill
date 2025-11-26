package dto

import (
	"errors"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseGetUserByID struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
}

func GetUserIDFromParam(ctx *gin.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return uuid.UUID{}, errors.New("invalid id")
	}
	return id, nil
}

func ToResponseGetUserByID(user *entity.User) *ResponseGetUserByID {
	return &ResponseGetUserByID{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		Age:         user.Age,
		Address:     user.Address,
	}
}
