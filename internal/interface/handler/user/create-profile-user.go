package handler

import (
	"net/http"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	dto "github.com/BAITOEYSRN/test-Technical-Skill/internal/interface/dto/user"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createProfileUserHandlerCfg struct {
	userUsecase domain.UserUsecase
}

func NewCreateProfileUserHandler(usecase domain.UserUsecase) *createProfileUserHandlerCfg {
	return &createProfileUserHandlerCfg{
		userUsecase: usecase,
	}
}

func (u *createProfileUserHandlerCfg) CreateProfileUserHandler(ctx *gin.Context) {
	req, err := new(dto.CreateProfileUserRequest).Validate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.CreateUser(ctx, entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		Address:     req.Address,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.ResponseJsonWithCode(ctx, http.StatusOK, uuid.New(), "success", "Create profile user success", dto.ToCreateProfileUserResponse(result))
}
