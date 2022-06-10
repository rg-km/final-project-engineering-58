package login_http_handler

import (
	"backend/domain/usecase"
	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/usecase/auth"

	"github.com/gorilla/mux"
)

type authHttpHandler struct {
	cfg         *config.Config
	authUsecase usecase.AuthUsecase
}

func AuthHttpHandler(r *mux.Router, repo repository.UserRepository, config *config.Config) {

	httpInteractor := auth.NewAuthInteractor(&repo, config)
	handler := &authHttpHandler{
		cfg:         config,
		authUsecase: httpInteractor,
	}

	r.HandleFunc("/api/auth/login", handler.Login).Methods("POST")
}
