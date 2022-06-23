package response

import (
	"backend/domain/entity"
	"backend/pkg/common"
	"time"
)

type PostResponse struct {
	ID       	common.ID	`json:"id"`
	Title	 	string    	`json:"title"`
	Description	string    	`json:"description"`
	Content		string    	`json:"content"`
	UrlVideo	string    	`json:"url_video"`
	CategoryID	string    	`json:"category_id"`
	UserID		string    	`json:"user_id"`
	ParentID	string    	`json:"parent_id"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
}

func MapPostDomainToResponse(category *entity.Post) PostResponse {
	return PostResponse{
		ID:              category.ID,
		Title:           category.Title,
		Description:     category.Description,
		Content:         category.Content,
		UrlVideo:        category.UrlVideo,
		CategoryID:      category.CategoryID,
		UserID:          category.UserID,
		ParentID:        category.ParentID,
		CreatedAt:       category.CreatedAt,
		UpdatedAt:       category.UpdatedAt,
	}
}

func MapPostListDomainToResponse(category []*entity.Post) []PostResponse {
	Posts := make([]PostResponse, 0)
	for _, Post := range category {
		Posts = append(Posts, MapPostDomainToResponse(Post))
	}
	return Posts
}