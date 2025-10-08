package handler

import (
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/internal/user/usecase"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
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
	params := request.GetPaginationParams(ctx)
	result, err := h.Usecase.FindAll(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}

	data, err := h.Usecase.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, response.NewRegularResponse(data, "User created successfully"))
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var req model.UpdateUserDto
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}

	data, err := h.Usecase.Update(ctx, id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.NewRegularResponse(data, "User updated successfully"))
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Usecase.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.NewRegularResponse(id, "User deleted successfully"))
}
