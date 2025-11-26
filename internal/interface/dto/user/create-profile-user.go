package dto

import (
	"errors"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateProfileUserRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type CreateProfileUserResponse struct {
	ID uuid.UUID `json:"id"`
}

func (r *CreateProfileUserRequest) Validate(ctx *gin.Context) (*CreateProfileUserRequest, error) {
	req := new(CreateProfileUserRequest)

	if err := ctx.BindJSON(&req); err != nil {
		return nil, err
	}
	if err := utils.Validator.Struct(req); err != nil {
		return nil, errors.New(err.Error())
	}
	return req, nil
}

func ToCreateProfileUserResponse(user *entity.User) *CreateProfileUserResponse {
	return &CreateProfileUserResponse{
		ID: user.ID,
	}
}
