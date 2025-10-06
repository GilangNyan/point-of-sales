package route

import (
	"gilangnyan/point-of-sales/internal/user/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	Handler handler.UserHandler
}

func NewUserRoutes(handler handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		Handler: handler,
	}
}

func (r *UserRoutes) Route(rg *gin.RouterGroup) {
	userHandler := rg.Group("/v1/users")

	userHandler.Use()
	{
		userHandler.GET("/", r.Handler.GetAllUsers)
		userHandler.POST("/", r.Handler.CreateUser)
		userHandler.PUT("/:id", r.Handler.UpdateUser)
		userHandler.DELETE("/:id", r.Handler.DeleteUser)
	}
}
