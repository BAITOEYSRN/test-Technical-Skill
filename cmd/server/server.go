package server

import (
	"fmt"
	"log"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/config"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/repository/db"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg"
	"github.com/gin-gonic/gin"
)

func Server(cfg *config.Config) *gin.Engine {
	dbs, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	fmt.Println("DB", dbs)
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

	router := gin.New()
	router.Use(pkg.Logging())
	router.Use(pkg.CORSMiddleware())
	router.Use(pkg.Recovery())
	router.Run(":8080")

	return nil
}
