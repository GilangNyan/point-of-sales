package server

import (
	"fmt"
	"gilangnyan/point-of-sales/internal/config"
	userDI "gilangnyan/point-of-sales/internal/user/di"
	"gilangnyan/point-of-sales/package/database"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	app  *gin.Engine
	db   database.Database
	conf *config.Config
}

func NewGinServer(conf *config.Config, db database.Database) *GinServer {
	ginApp := gin.Default()

	return &GinServer{
		app:  ginApp,
		db:   db,
		conf: conf,
	}
}

// setupRoutes initializes all application routes
func (s *GinServer) setupRoutes() {
	// Create API v1 group
	apiV1 := s.app.Group("/api")

	// Initialize User module with DI
	userModule := userDI.NewUserModule(s.db.GetDB())
	userModule.RegisterRoutes(apiV1)

	// TODO: Add other modules here
	// productModule := productDI.NewProductModule(s.db.GetDB())
	// productModule.RegisterRoutes(apiV1)
}

func (s *GinServer) Start() {
	s.app.Use(gin.Recovery())
	s.app.Use(gin.Logger())

	// Health check endpoint
	s.app.GET("/v1/health", func(c *gin.Context) {
		// Check database health
		dbHealth, err := s.db.Health()
		if err != nil {
			c.JSON(500, gin.H{
				"status":   "unhealthy",
				"message":  "Database connection failed",
				"database": dbHealth,
				"error":    err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status":   "healthy",
			"message":  "Service is running",
			"database": dbHealth,
		})
	})

	// Setup all application routes
	s.setupRoutes()

	serverUrl := fmt.Sprintf("%s:%s", s.conf.Server.Host, s.conf.Server.Port)
	s.app.Run(serverUrl)
}
