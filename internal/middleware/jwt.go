package middleware

import (
	"gilangnyan/point-of-sales/package/jwt"
	"gilangnyan/point-of-sales/package/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTUserMiddleware struct {
	JWTService jwt.JWTService
}

func NewJWTUserMiddleware(jwtService jwt.JWTService) *JWTUserMiddleware {
	return &JWTUserMiddleware{
		JWTService: jwtService,
	}
}

func (m *JWTUserMiddleware) Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse("UNAUTHORIZED", "Authorization header is required"))
			return
		}

		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse("UNAUTHORIZED", "Invalid Authorization header format"))
			return
		}

		claims, err := m.JWTService.ValidateToken(tokenParts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse("UNAUTHORIZED", err.Error()))
			return
		}

		// TODO: Check user session in DB
		// currentTokenSession, err := m.AuthSession.GetUserSession()

		ctx.Set("userID", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("email", claims.Email)
		ctx.Set("roles", claims.Roles)

		ctx.Next()
	}
}
