package category_http_handler

import (
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (x *categoryHttpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, role, errJwt := utils.DecodeJwtToken(r)
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	if role != "admin" {
		utils.RespondWithError(w, http.StatusUnauthorized, []error{errors.New("not authorization")})
		return
	}

	err := x.categoryUsecase.Delete(context.Background(), id)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}