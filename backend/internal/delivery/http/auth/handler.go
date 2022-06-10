package login_http_handler

import (
	"backend/domain/usecase"
	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/usecase/user"

	"github.com/gorilla/mux"
)

type loginHttpHandler struct {
	cfg         *config.Config
	userUsecase usecase.UserUsecase
}

func LoginHttpHandler(r *mux.Router, repo repository.UserRepository, config *config.Config) {

	httpInteractor := user.LoginInteractor(&repo, config)
	handler := &loginHttpHandler{
		cfg:         config,
		userUsecase: httpInteractor,
	}

	r.HandleFunc("/api/login", handler.Login).Methods("POST")
}
