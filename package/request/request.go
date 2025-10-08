package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	SortBy   string `json:"sortBy"`
	SortDir  string `json:"sortDir"`
}

func GetPaginationParams(ctx *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	sortBy := ctx.DefaultQuery("sortBy", "id")
	sortDir := ctx.DefaultQuery("sortDir", "ASC")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	if sortDir != "ASC" && sortDir != "DESC" {
		sortDir = "ASC"
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
		SortDir:  sortDir,
	}
}

func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationParams) GetLimit() int {
	return p.PageSize
}
