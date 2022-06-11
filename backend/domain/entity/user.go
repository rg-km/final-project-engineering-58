package entity

import (
	"backend/pkg/common"
	"backend/pkg/utils"
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"
)

type User struct {
	ID        common.ID
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDto struct {
	Name     string
	Email    string
	Password string
	Role     string
}

func NewUser(payload UserDto) *User {
	create := User{
		ID:        common.NewID(),
		Name:      payload.Name,
		Email:     payload.Email,
		Role:      payload.Role,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return &create
}

func (payload *User) Validate() *multierror.Error {
	var multilerr *multierror.Error

	if payload.Name == "" {
		multilerr = multierror.Append(multilerr, errors.New("name required"))
	}

	if payload.Email == "" {
		multilerr = multierror.Append(multilerr, errors.New("email required"))
	}

	if payload.Password == "" {
		multilerr = multierror.Append(multilerr, errors.New("password required"))
	}

	if payload.Role == "" {
		multilerr = multierror.Append(multilerr, errors.New("role required"))
	}

	return multilerr
}

func (payload *User) SetRole(role string) {
	payload.Role = role
}

func (payload *User) SetPasswordHash(password string) {
	pwd, _ := utils.HashPassword(password)
	payload.Password = pwd
}
