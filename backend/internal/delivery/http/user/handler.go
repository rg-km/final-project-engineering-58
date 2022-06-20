package user_http_handler

import (
	"backend/domain/usecase"
	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/usecase/user"

	"github.com/gorilla/mux"
)

type userHttpHandler struct {
	cfg         *config.Config
	userUsecase usecase.UserUsecase
}

func UserHttpHandler(r *mux.Router, repo repository.UserRepository, config *config.Config) {
	httpInteractor := user.NewUserInteractor(&repo, config)
	handler := &userHttpHandler{
		cfg:         config,
		userUsecase: httpInteractor,
	}

	r.HandleFunc("/api/user", handler.Create).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/api/user/{id}", handler.Delete).Methods("DELETE")
	r.HandleFunc("/api/user", handler.GetAll).Methods("GET")
}
