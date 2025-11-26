package handler

import (
	"net/http"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain"
	dto "github.com/BAITOEYSRN/test-Technical-Skill/internal/interface/dto/user"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type getUserByIDHandler struct {
	userUsecase domain.UserUsecase
}

func NewGetUserByIDHandler(usecase domain.UserUsecase) *getUserByIDHandler {
	return &getUserByIDHandler{
		userUsecase: usecase,
	}
}

func (u *getUserByIDHandler) GetUserByIDHandler(ctx *gin.Context) {
	id, err := dto.GetUserIDFromParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userUsecase.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.ResponseJsonWithCode(ctx, http.StatusOK, uuid.New(), "success", "Get user by id success", dto.ToResponseGetUserByID(user))
}
