package api

import (
	"magic-bullet/backend/internal/api/handlers"
	"magic-bullet/backend/internal/config"
	"magic-bullet/backend/internal/database"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

const basePath = "/api/v1"

type App struct {
	Engine *gin.Engine
	API    huma.API
	Config *config.Config
	Database database.Database
}

func NewApp(config *config.Config, databsae database.Database) *App {
	// Set Gin mode based on environment variable, defaulting to release mode
	gin.SetMode(envOrDefault("GIN_MODE", gin.ReleaseMode))

	// Create a new Gin router with logging and recovery middleware
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Configure Huma API with OpenAPI metadata
	apiCFG := huma.DefaultConfig("Magic Bullet API", "1.0.0")
	apiCFG.Servers = []*huma.Server{{URL: basePath}}

	// Create Huma API instance and register it with the Gin router
	v1 := router.Group(basePath)
	api := humagin.NewWithGroup(router, v1, apiCFG)

	// Initialize Handler and register routes
	handlers := handlers.NewHandler(api)
	handlers.RegisterRoutes()

	return &App{
		Engine: router,
		API:    api,
		Config: config,
	}
}

func HTTPAddr() string {
	return envOrDefault("HTTP_ADDR", ":8080")
}

func envOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	return fallback
}
