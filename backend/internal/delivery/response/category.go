package response

import (
	"backend/domain/entity"
	"backend/pkg/common"
	"time"
)

type CategoryResponse struct {
	ID       	common.ID	`json:"id"`
	Name     	string    	`json:"name"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
}

func MapCategoryDomainToResponse(category *entity.Category) CategoryResponse {
	return CategoryResponse{
		ID:              category.ID,
		Name:            category.Name,
		CreatedAt:       category.CreatedAt,
		UpdatedAt:       category.UpdatedAt,
	}
}

func MapCategoryListDomainToResponse(category []*entity.Category) []CategoryResponse {
	Categorys := make([]CategoryResponse, 0)
	for _, Category := range category {
		Categorys = append(Categorys, MapCategoryDomainToResponse(Category))
	}
	return Categorys
}