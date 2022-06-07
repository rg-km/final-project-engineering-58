package response

import (
	"backend/pkg/common"
	"backend/domain/entity"
)

type UserResponse struct {
	ID			common.ID `json:"id"`
	Name		string    `json:"name"`
	Email		string    `json:"email"`
	CreatedAt	string    `json:"created_at"`
	UpdatedAt	string    `json:"updated_at"`
}

func MapUserDomainToResponse(user *entity.User) UserResponse {
	return UserResponse{
		ID:			user.ID,
		Name:		user.Name,
		Email:		user.Email,
		CreatedAt:	user.CreatedAt.String(),
		UpdatedAt:	user.UpdatedAt.String(),
	}
}