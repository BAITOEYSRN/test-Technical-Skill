package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BAITOEYSRN/test-Technical-Skill/cmd/server"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	s := server.Server(cfg)
	fmt.Println("Server started", s)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-shutdownCtx.Done()
	fmt.Println("Server shutdown complete")

}
