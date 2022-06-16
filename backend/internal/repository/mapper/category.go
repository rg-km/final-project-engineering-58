package mapper

import (
	"backend/domain/entity"
	"backend/internal/repository/models"
	"backend/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToModelCategory(dt *entity.Category) *models.Category {
	return &models.Category{
		ID:        	dt.ID.String(),
		Name: 		dt.Name,
		CreatedAt:	dt.CreatedAt,
		UpdatedAt:	dt.UpdatedAt,
	}
}

func ToDomainCategory(m *models.Category) *entity.Category {
	id, _ := common.StringToID(m.ID)

	return &entity.Category{
		ID:			id,
		Name: 	m.Name,
		CreatedAt:	m.CreatedAt,
		UpdatedAt:	m.UpdatedAt,
	}
}

func ToDbqStructCategory(domain *entity.Category) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqCategory(domain))
	return
}

func DataDbqCategory(domain *entity.Category) []interface{} {
	return dbq.Struct(ToModelCategory(domain))
}

func ToListDomainCategory(dto []*models.Category) []*entity.Category {
	Categorys := make([]*entity.Category, 0)
	for _, Category := range dto {
		Categorys = append(Categorys, ToDomainCategory(Category))
	}
	return Categorys
}