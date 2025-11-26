package routes

import (
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/repository"
	handler "github.com/BAITOEYSRN/test-Technical-Skill/internal/interface/handler/user"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfigUserRoutesCfg struct {
	DB  *gorm.DB
	App *gin.Engine
}

func (c *ConfigUserRoutesCfg) NewConfigUserRoutes() {
	repo := repository.NewUserRepository(c.DB)
	userUsecase := usecase.NewUserRepository(repo)

	user := c.App.Group("/api/v1/users")
	user.POST("", handler.NewCreateProfileUserHandler(userUsecase).CreateProfileUserHandler)
	user.GET("/list", handler.NewGetListProfileUserHandler(userUsecase).GetListProfileUserHandler)
	user.GET("/:id", handler.NewGetUserByIDHandler(userUsecase).GetUserByIDHandler)
}
