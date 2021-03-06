package user_http_handler

import (
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (x *userHttpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := x.userUsecase.Delete(context.Background(), id)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}
