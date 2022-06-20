package response

import (
	"backend/domain/entity"
	"backend/pkg/common"
)

type UserResponse struct {
	ID        common.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func MapUserDomainToResponse(user *entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func MapUserListDomainToResponse(user []*entity.User) []UserResponse {
	Users := make([]UserResponse, 0)
	for _, User := range user {
		Users = append(Users, MapUserDomainToResponse(User))
	}
	return Users
}
