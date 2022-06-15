package category

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *categoryInteractor) Update(ctx context.Context, payload entity.CategoryDto, id string) (*entity.Category, *exceptions.CustomError) {
	var multilerr *multierror.Error

	category, errFind := x.repo.FindByID(ctx, id)
	if errFind != nil {
		multilerr = multierror.Append(multilerr, errFind)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	update := entity.NewCategory(payload)

	if err := update.Validate(); err != nil {
		return nil, &exceptions.CustomError{
			Errors: err,
			Status: exceptions.ERRDOMAIN,
		}
	}

	err := x.repo.Update(ctx, update, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	category.Name = update.Name

	return category, nil
}