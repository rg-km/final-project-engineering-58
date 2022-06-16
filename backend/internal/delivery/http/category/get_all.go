package category_http_handler

import (
	"backend/internal/delivery/response"
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"net/http"
)

func (x *categoryHttpHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := x.categoryUsecase.GetAll(context.Background())

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapCategoryListDomainToResponse(categories))
}