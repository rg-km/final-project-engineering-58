package entity

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type Auth struct {
	Email   string
	Password string
}

func (payload *Auth) Validate() *multierror.Error {
	var multilerr *multierror.Error

	if payload.Email == "" {
		multilerr = multierror.Append(multilerr, errors.New("email required"))
	}

	if payload.Password == "" {
		multilerr = multierror.Append(multilerr, errors.New("password required"))
	}

	return multilerr
}