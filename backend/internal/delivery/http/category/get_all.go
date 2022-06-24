package category_http_handler

import (
	"backend/internal/delivery/response"
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"errors"
	"net/http"
)

func (x *categoryHttpHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	_, role, errJwt := utils.DecodeJwtToken(r)
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	if role != "admin" {
		utils.RespondWithError(w, http.StatusUnauthorized, []error{errors.New("not authorization")})
		return
	}

	categories, err := x.categoryUsecase.GetAll(context.Background())

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapCategoryListDomainToResponse(categories))
}