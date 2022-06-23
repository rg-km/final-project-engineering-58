package mapper

import (
	"backend/domain/entity"
	"backend/internal/repository/models"
	"backend/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToModelPost(dt *entity.Post) *models.Post {
	return &models.Post{
		ID:        		dt.ID.String(),
		Title:    		dt.Title,
		Description: 	dt.Description,
		Content:  		dt.Content,
		UrlVideo:  		dt.UrlVideo,
		CategoryID:		dt.CategoryID,
		UserID:    		dt.UserID,
		ParentID:  		dt.ParentID,
		CreatedAt:		dt.CreatedAt,
		UpdatedAt:		dt.UpdatedAt,
	}
}

func ToDomainPost(m *models.Post) *entity.Post {
	id, _ := common.StringToID(m.ID)

	return &entity.Post{
		ID:				id,
		Title:    		m.Title,
		Description: 	m.Description,
		Content:  		m.Content,
		UrlVideo:  		m.UrlVideo,
		CategoryID:		m.CategoryID,
		UserID:    		m.UserID,
		ParentID:  		m.ParentID,
		CreatedAt:		m.CreatedAt,
		UpdatedAt:		m.UpdatedAt,
	}
}

func ToDbqStructPost(domain *entity.Post) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqPost(domain))
	return
}

func DataDbqPost(domain *entity.Post) []interface{} {
	return dbq.Struct(ToModelPost(domain))
}

func ToListDomainPost(dto []*models.Post) []*entity.Post {
	Posts := make([]*entity.Post, 0)
	for _, Post := range dto {
		Posts = append(Posts, ToDomainPost(Post))
	}
	return Posts
}