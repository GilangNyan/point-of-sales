package handler

import (
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		Usecase: usecase,
	}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	data, err := h.Usecase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	data, err := h.Usecase.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": data})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var req model.UpdateUserDto
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	data, err := h.Usecase.Update(ctx, id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Usecase.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
