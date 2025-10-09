package route

import (
	"gilangnyan/point-of-sales/internal/role/handler"

	"github.com/gin-gonic/gin"
)

type RoleRoutes struct {
	Handler handler.RoleHandler
}

func NewRoleRoutes(handler handler.RoleHandler) *RoleRoutes {
	return &RoleRoutes{
		Handler: handler,
	}
}

func (r *RoleRoutes) Route(rg *gin.RouterGroup) {
	roleHandler := rg.Group("/v1/roles")

	roleHandler.Use()
	{
		roleHandler.GET("/", r.Handler.GetAllRoles)
		roleHandler.POST("/", r.Handler.CreateRoles)
		roleHandler.PUT("/:id", r.Handler.UpdateRoles)
		roleHandler.DELETE("/:id", r.Handler.DeleteRoles)
	}
}
