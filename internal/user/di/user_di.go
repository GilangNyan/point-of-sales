package di

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/handler"
	"gilangnyan/point-of-sales/internal/user/repository"
	"gilangnyan/point-of-sales/internal/user/route"
	"gilangnyan/point-of-sales/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// UserModule represents the complete user module with all dependencies
type UserModule struct {
	Repository repository.UserRepository
	Usecase    usecase.UserUsecase
	Handler    *handler.UserHandler
	Routes     *route.UserRoutes
}

// NewUserModule creates a new user module with all dependencies injected
func NewUserModule(db *sql.DB) *UserModule {
	// Build dependency chain
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(uc)
	routes := route.NewUserRoutes(*handler)

	return &UserModule{
		Repository: repo,
		Usecase:    uc,
		Handler:    handler,
		Routes:     routes,
	}
}

// RegisterRoutes registers all user routes to the router group
func (um *UserModule) RegisterRoutes(rg *gin.RouterGroup) {
	um.Routes.Route(rg)
}

// Provider functions for individual components (if needed separately)

// ProvideUserRepository creates user repository
func ProvideUserRepository(db *sql.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

// ProvideUserUsecase creates user usecase
func ProvideUserUsecase(repo repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(repo)
}

// ProvideUserHandler creates user handler
func ProvideUserHandler(uc usecase.UserUsecase) *handler.UserHandler {
	return handler.NewUserHandler(uc)
}

// ProvideUserRoutes creates user routes
func ProvideUserRoutes(handler *handler.UserHandler) *route.UserRoutes {
	return route.NewUserRoutes(*handler)
}
