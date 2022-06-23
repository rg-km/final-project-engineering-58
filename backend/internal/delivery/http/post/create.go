package http_handler

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

func (x *postHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload request.CreatePostPayload
		ctx = context.Background()
	)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		msg := errors.New("payload required")
		utils.RespondWithError(w, http.StatusBadRequest, []error{msg})
		return
	}

	userId, role, errJwt := utils.DecodeJwtToken(r)
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	if role != "admin" && role != "constributor" {
		utils.RespondWithError(w, http.StatusUnauthorized, []error{errors.New("not authorization")})
		return
	}

	create := entity.PostDto{
		Title: payload.Title,
		Content: payload.Content,
		Description: payload.Description,
		UrlVideo: payload.UrlVideo,
		CategoryID: payload.CategoryID,
		UserID: userId,
		ParentID: payload.ParentID,
	}

	category, err := x.postUsecase.Create(ctx, create)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapPostDomainToResponse(category))
}