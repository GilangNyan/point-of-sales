package route

import (
	"gilangnyan/point-of-sales/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	handler *handler.AuthHandler
}

func NewAuthRoutes(handler *handler.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
	}
}

func (r *AuthRoutes) Route(rg *gin.RouterGroup) {
	authHandler := rg.Group("/v1/auth")

	authHandler.Use()
	{
		authHandler.POST("/login", r.handler.Login)
	}
}
