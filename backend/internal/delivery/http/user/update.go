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

	"github.com/gorilla/mux"
)

func (x *userHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		payload request.CreateUserPayload
		ctx     = context.Background()
		vars    = mux.Vars(r)
		id      = vars["id"]
	)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		msg := errors.New("payload required")
		utils.RespondWithError(w, http.StatusBadRequest, []error{msg})
		return
	}

	update := entity.UserDto{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	user, err := x.userUsecase.Update(ctx, update, id)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserDomainToResponse(user))
}
