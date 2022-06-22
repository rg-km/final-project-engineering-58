package http_handler

import (
	"backend/domain/usecase"
	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/usecase/post"

	"github.com/gorilla/mux"
)

type postHttpHandler struct {
	cfg         *config.Config
	postUsecase usecase.PostUsecase
}

func PostHttpHandler(r *mux.Router, repo repository.PostRepository, config *config.Config) {
	httpInteractor := post.NewPostInteractor(&repo, config)
	handler := &postHttpHandler{
		cfg:         config,
		postUsecase: httpInteractor,
	}

	r.HandleFunc("/api/posts", handler.GetAll).Methods("GET")
	r.HandleFunc("/api/post", handler.Create).Methods("POST")
	r.HandleFunc("/api/post/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/api/post/{id}", handler.Delete).Methods("DELETE")
}