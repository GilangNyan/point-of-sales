package handler

import (
	"gilangnyan/point-of-sales/internal/role/model"
	"gilangnyan/point-of-sales/internal/role/usecase"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	Usecase usecase.RoleUsecase
}

func NewRoleHandler(usecase usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		Usecase: usecase,
	}
}

func (h *RoleHandler) GetAllRoles(ctx *gin.Context) {
	params := request.GetPaginationParams(ctx)
	result, err := h.Usecase.FindAll(ctx, params)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *RoleHandler) CreateRoles(ctx *gin.Context) {
	var req model.CreateRoleDto

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

	ctx.JSON(http.StatusCreated, response.NewRegularResponse(data, "Role created successfully"))
}

func (h *RoleHandler) UpdateRoles(ctx *gin.Context) {
	var req model.UpdateRoleDto
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

	ctx.JSON(http.StatusOK, response.NewRegularResponse(data, "Role updated successfully"))
}

func (h *RoleHandler) DeleteRoles(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Usecase.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("INTERNAL_SERVER_ERROR", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.NewRegularResponse(id, "Role deleted successfully"))
}
