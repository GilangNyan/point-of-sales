package di

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/middleware"
	"gilangnyan/point-of-sales/internal/user/handler"
	"gilangnyan/point-of-sales/internal/user/repository"
	"gilangnyan/point-of-sales/internal/user/route"
	"gilangnyan/point-of-sales/internal/user/usecase"
	userRoleRepo "gilangnyan/point-of-sales/internal/user_role/repository"
	"gilangnyan/point-of-sales/package/transaction"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
	Repository   repository.UserRepository
	Repository2  repository.UserProfileRepository
	UserRoleRepo userRoleRepo.UserRoleRepository
	Usecase      usecase.UserUsecase
	Handler      *handler.UserHandler
	Routes       *route.UserRoutes
}

// NewUserModule creates a new user module with all dependencies injected
func NewUserModule(db *sql.DB, middleware *middleware.JWTUserMiddleware) *UserModule {
	// Build dependency chain
	repo := repository.NewUserRepository(db)
	repo2 := repository.NewUserProfileRepository(db)
	userRoleRepo := userRoleRepo.NewUserRoleRepository(db)
	tx := transaction.NewTransactionManager(db)
	uc := usecase.NewUserUsecase(repo, repo2, userRoleRepo, tx)
	handler := handler.NewUserHandler(uc)
	routes := route.NewUserRoutes(*handler, middleware)

	return &UserModule{
		Repository:   repo,
		Repository2:  repo2,
		UserRoleRepo: userRoleRepo,
		Usecase:      uc,
		Handler:      handler,
		Routes:       routes,
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
func ProvideUserUsecase(repo repository.UserRepository, repo2 repository.UserProfileRepository, userRoleRepo userRoleRepo.UserRoleRepository, tx transaction.TransactionManager) usecase.UserUsecase {
	return usecase.NewUserUsecase(repo, repo2, userRoleRepo, tx)
}

// ProvideUserHandler creates user handler
func ProvideUserHandler(uc usecase.UserUsecase) *handler.UserHandler {
	return handler.NewUserHandler(uc)
}

// ProvideUserRoutes creates user routes
func ProvideUserRoutes(handler *handler.UserHandler, middleware *middleware.JWTUserMiddleware) *route.UserRoutes {
	return route.NewUserRoutes(*handler, middleware)
}
