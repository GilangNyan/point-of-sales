package model

import "gilangnyan/point-of-sales/package/template"

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	template.Base
}
