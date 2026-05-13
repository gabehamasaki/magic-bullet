package api

import (
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

const basePath = "/api/v1"

type App struct {
	Engine *gin.Engine
	API    huma.API
}

func NewApp() *App {
	gin.SetMode(envOrDefault("GIN_MODE", gin.ReleaseMode))

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	registerHealthRoute(router)

	config := huma.DefaultConfig("Magic Bullet API", "1.0.0")
	config.Servers = []*huma.Server{{URL: basePath}}

	v1 := router.Group(basePath)
	api := humagin.NewWithGroup(router, v1, config)

	registerUserRoutes(api)

	return &App{
		Engine: router,
		API:    api,
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
