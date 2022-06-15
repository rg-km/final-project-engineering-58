package entity

import (
	"backend/pkg/common"
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Category struct {
	ID        common.ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryDto struct {
	Name string
}

func NewCategory(payload CategoryDto) *Category {
	create := Category{
		ID:        common.NewID(),
		Name:      payload.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return &create
}

func (x *Category) Validate() *multierror.Error {
	var multilerr *multierror.Error

	if x.Name == "" {
		multilerr = multierror.Append(multilerr, errors.New("name required"))
	}

	return multilerr
}