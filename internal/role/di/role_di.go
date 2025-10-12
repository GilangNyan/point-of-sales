package di

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/middleware"
	"gilangnyan/point-of-sales/internal/role/handler"
	"gilangnyan/point-of-sales/internal/role/repository"
	"gilangnyan/point-of-sales/internal/role/route"
	"gilangnyan/point-of-sales/internal/role/usecase"

	"github.com/gin-gonic/gin"
)

type RoleModule struct {
	Repository repository.RoleRepository
	Usecase    usecase.RoleUsecase
	Handler    *handler.RoleHandler
	Routes     *route.RoleRoutes
}

func NewRoleModule(db *sql.DB, middleware *middleware.JWTUserMiddleware) *RoleModule {
	repo := repository.NewRoleRepository(db)
	uc := usecase.NewRoleUsecase(repo)
	handler := handler.NewRoleHandler(uc)
	routes := route.NewRoleRoutes(*handler, middleware)

	return &RoleModule{
		Repository: repo,
		Usecase:    uc,
		Handler:    handler,
		Routes:     routes,
	}
}

func (rm *RoleModule) RegisterRoutes(rg *gin.RouterGroup) {
	rm.Routes.Route(rg)
}
