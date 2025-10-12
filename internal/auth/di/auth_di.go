package di

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/auth/handler"
	"gilangnyan/point-of-sales/internal/auth/repository"
	"gilangnyan/point-of-sales/internal/auth/route"
	"gilangnyan/point-of-sales/internal/auth/usecase"
	userRoleRepo "gilangnyan/point-of-sales/internal/user_role/repository"
	"gilangnyan/point-of-sales/package/jwt"

	"github.com/gin-gonic/gin"
)

type AuthModule struct {
	Repository   repository.AuthRepository
	UserRoleRepo userRoleRepo.UserRoleRepository
	Usecase      usecase.AuthUsecase
	Handler      *handler.AuthHandler
	Routes       *route.AuthRoutes
}

func NewAuthModule(db *sql.DB, jwtService jwt.JWTService) *AuthModule {
	repo := repository.NewAuthRepository(db)
	userRoleRepo := userRoleRepo.NewUserRoleRepository(db)
	uc := usecase.NewAuthUsecase(repo, userRoleRepo, jwtService)
	handler := handler.NewAuthHandler(uc)
	routes := route.NewAuthRoutes(handler)

	return &AuthModule{
		Repository:   repo,
		UserRoleRepo: userRoleRepo,
		Usecase:      uc,
		Handler:      handler,
		Routes:       routes,
	}
}

func (am *AuthModule) RegisterRoutes(rg *gin.RouterGroup) {
	am.Routes.Route(rg)
}
