package server

import (
	"fmt"
	authDI "gilangnyan/point-of-sales/internal/auth/di"
	"gilangnyan/point-of-sales/internal/config"
	"gilangnyan/point-of-sales/internal/middleware"
	roleDI "gilangnyan/point-of-sales/internal/role/di"
	userDI "gilangnyan/point-of-sales/internal/user/di"
	"gilangnyan/point-of-sales/package/database"
	"gilangnyan/point-of-sales/package/jwt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	app           *gin.Engine
	db            database.Database
	conf          *config.Config
	jwtService    jwt.JWTService
	jwtMiddleware *middleware.JWTUserMiddleware
}

func NewGinServer(conf *config.Config, db database.Database) *GinServer {
	ginApp := gin.Default()

	// Initiaize JWT Service
	jwtSecret := "824a9be2e3d42453299d4b777d52a5b3b333a4afefffb87e0d96dce505873095"
	jwtIssuer := "gilangnyan"

	expHours, _ := strconv.Atoi("1")
	refreshExpHours, _ := strconv.Atoi("6")

	jwtService := jwt.NewJWTService(
		jwtSecret,
		jwtIssuer,
		time.Duration(expHours)*time.Hour,
		time.Duration(refreshExpHours)*time.Hour,
	)

	// Initialize Middlewares
	jwtMiddleware := middleware.NewJWTUserMiddleware(jwtService)

	return &GinServer{
		app:           ginApp,
		db:            db,
		conf:          conf,
		jwtService:    jwtService,
		jwtMiddleware: jwtMiddleware,
	}
}

// setupRoutes initializes all application routes
func (s *GinServer) setupRoutes() {
	// Create API v1 group
	apiV1 := s.app.Group("/api")

	// Auth Module Routes
	authModule := authDI.NewAuthModule(s.db.GetDB(), s.jwtService)
	authModule.RegisterRoutes(apiV1)

	// User Module Routes
	userModule := userDI.NewUserModule(s.db.GetDB(), s.jwtMiddleware)
	userModule.RegisterRoutes(apiV1)

	// Role Module Routes
	roleModule := roleDI.NewRoleModule(s.db.GetDB())
	roleModule.RegisterRoutes(apiV1)
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
