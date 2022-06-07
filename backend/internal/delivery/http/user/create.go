package user_http_handler

import (
	"backend/domain/entity"
	"backend/internal/delivery/request"
	"backend/internal/delivery/response"
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (x userHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload request.CreateUserPayload
		ctx = context.Background()
	)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		msg := errors.New("payload required")
		utils.RespondWithError(w, http.StatusBadRequest, []error{msg})
		return
	}

	create := entity.UserDto{
		Name: payload.Name,
		Email: payload.Email,
	}

	user, err := x.userUsecase.Create(ctx, create)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}
	
	utils.RespondWithJSON(w, http.StatusCreated, response.MapUserDomainToResponse(user))
}