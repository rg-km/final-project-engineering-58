package auth_http_handler

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
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

func (x authHttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		payload request.CreateAuthPayload
		ctx     = context.Background()
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		msg := errors.New("payload required")
		utils.RespondWithError(w, http.StatusBadRequest, []error{msg})
		return
	}

	create := entity.Auth{
		Email:    payload.Email,
		Password: payload.Password,
	}

	user, err := x.authUsecase.Login(ctx, create)

	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, errTkn := token.SignedString(jwtKey)
	if errTkn != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, []error{errTkn})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   strToken,
		Expires: expirationTime,
	})

	dataResp := response.MapUserDomainToResponse(user)
	response, _ := json.Marshal(jsonResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    dataResp,
		Token:   strToken,
		Meta:    nil,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
