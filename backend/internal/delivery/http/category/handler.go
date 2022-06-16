package category_http_handler

import (
	"backend/domain/usecase"
	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/usecase/category"

	"github.com/gorilla/mux"
)

type categoryHttpHandler struct {
	cfg         *config.Config
	categoryUsecase usecase.CategoryUsecase
}

func CategoryHttpHandler(r *mux.Router, repo repository.CategoryRepository, config *config.Config) {
	httpInteractor := category.NewCategoryInteractor(&repo, config)
	handler := &categoryHttpHandler{
		cfg:         config,
		categoryUsecase: httpInteractor,
	}

	r.HandleFunc("/api/category", handler.Create).Methods("POST")
	r.HandleFunc("/api/category", handler.GetAll).Methods("GET")
	r.HandleFunc("/api/category/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/api/category/{id}", handler.Delete).Methods("DELETE")
}