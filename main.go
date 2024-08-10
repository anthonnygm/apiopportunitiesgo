package main

import (
	"github.com/anthonnygm/apiopportunitiesgo/config"
	"github.com/anthonnygm/apiopportunitiesgo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.BuildLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
