package response

import "math"

type RegularResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type PaginationResponse[T any] struct {
	Data      []T   `json:"data"`
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func NewRegularResponse[T any](data T, message string) *RegularResponse[T] {
	return &RegularResponse[T]{
		Data:    data,
		Message: message,
		Success: true,
	}
}

func NewPaginationResponse[T any](data []T, total int64, page, pageSize int) *PaginationResponse[T] {
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))

	return &PaginationResponse[T]{
		Data:      data,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}
}

func NewErrorResponse(err string, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   err,
		Message: message,
		Success: false,
	}
}
