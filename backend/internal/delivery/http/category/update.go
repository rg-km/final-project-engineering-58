package category_http_handler

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

func (x *categoryHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		payload request.CreateCategoryPayload
		ctx = context.Background()
		vars = mux.Vars(r)
		id = vars["id"]
	)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		msg := errors.New("payload required")
		utils.RespondWithError(w, http.StatusBadRequest, []error{msg})
		return
	}

	update := entity.CategoryDto{
		Name: payload.Name,
	}

	category, err := x.categoryUsecase.Update(ctx, update, id)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapCategoryDomainToResponse(category))
}