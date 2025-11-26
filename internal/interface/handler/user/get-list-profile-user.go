package handler

import (
	"net/http"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain"
	dto "github.com/BAITOEYSRN/test-Technical-Skill/internal/interface/dto/user"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type getListProfileUserHandler struct {
	userUsecase domain.UserUsecase
}

func NewGetListProfileUserHandler(usecase domain.UserUsecase) *getListProfileUserHandler {
	return &getListProfileUserHandler{
		userUsecase: usecase,
	}
}

func (u *getListProfileUserHandler) GetListProfileUserHandler(ctx *gin.Context) {
	getlistUser, err := u.userUsecase.GetListUsers(ctx)
	if err != nil {
		response.ResponseErrorJsonWithCode(ctx, err)
		return
	}
	response.ResponseJsonWithCode(ctx, http.StatusOK, uuid.New(), "success", "Get list profile user success", dto.ToResponseListUser(getlistUser))
}
