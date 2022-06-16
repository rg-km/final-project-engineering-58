package auth_http_handler

import (
	"backend/internal/delivery/response"
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"net/http"
)

func (x authHttpHandler) Profile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	userId, _, errJwt := utils.DecodeJwtToken(r)
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	user, err := x.authUsecase.Profile(ctx, userId)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}
	
	utils.RespondWithJSON(w, http.StatusOK, response.MapUserDomainToResponse(user))
}
