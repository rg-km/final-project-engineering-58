package mapper

import (
	"backend/domain/entity"
	"backend/internal/repository/models"
	"backend/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToModelUser(dt *entity.User) *models.User {
	return &models.User{
		ID:        	dt.ID.String(),
		Name: 		dt.Name,
		Role:		dt.Role,
		Password: 	dt.Password,
		Email:		dt.Email,
		CreatedAt:	dt.CreatedAt,
		UpdatedAt:	dt.UpdatedAt,
	}
}

func ToDomainUser(m *models.User) *entity.User {
	id, _ := common.StringToID(m.ID)

	return &entity.User{
		ID:			id,
		Name: 	m.Name,
		Email:	m.Email,
		Password:	m.Password,
		Role:	m.Role,
		CreatedAt:	m.CreatedAt,
		UpdatedAt:	m.UpdatedAt,
	}
}

func ToDbqStructUser(domain *entity.User) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqUser(domain))
	return
}

func DataDbqUser(domain *entity.User) []interface{} {
	return dbq.Struct(ToModelUser(domain))
}

func ToListDomainUser(dto []*models.User) []*entity.User {
	users := make([]*entity.User, 0)
	for _, user := range dto {
		users = append(users, ToDomainUser(user))
	}
	return users
}