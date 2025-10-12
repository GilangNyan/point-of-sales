package handler

import (
	"gilangnyan/point-of-sales/internal/auth/model"
	"gilangnyan/point-of-sales/internal/auth/usecase"
	"gilangnyan/point-of-sales/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req model.LoginDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}

	result, err := h.authUsecase.Login(ctx, req)
	if err != nil {
		if err.Error() == "Invalid email or password" {
			ctx.JSON(http.StatusUnauthorized, response.NewErrorResponse("UNAUTHORIZED", err.Error()))
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.NewRegularResponse(result, "Login successful"))
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}
