package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/config"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/db"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/interface/routes"
	middleware "github.com/BAITOEYSRN/test-Technical-Skill/pkg/middleware"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbConn, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	if cfg.MigrationDBAuto {
		migrations := []db.MigrationConfig{
			{
				Path: cfg.PathMigrations,
			},
		}
		err := db.MigrateDB(cfg, migrations)
		if err != nil {
			log.Fatal("Failed to migrate database: ", err)
		}
		fmt.Println("Migration database successful")
	}

	app := gin.New()

	app.GET("/health", func(c *gin.Context) {
		response.ResponseJsonWithCode(c, http.StatusOK, uuid.New(), "success", "OK", nil)
	})

	app.Use(middleware.Logging())
	app.Use(middleware.CORSMiddleware())

	userRoutes := routes.ConfigUserRoutesCfg{
		DB:  dbConn,
		App: app,
	}

	userRoutes.NewConfigUserRoutes()

	app.Use(middleware.Recovery())
	app.Run(cfg.PORT)

}
