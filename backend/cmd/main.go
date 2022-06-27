package main

import (
	"backend/internal/config"
	auth_handler "backend/internal/delivery/http/auth"
	category_handler "backend/internal/delivery/http/category"
	post_handler "backend/internal/delivery/http/post"
	user_handler "backend/internal/delivery/http/user"
	repositories "backend/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	cfg = config.InitConfig()
	db  = config.InitDatabase()

	userRepository     = repositories.NewUserRepository(db)
	categoryRepository = repositories.NewCategoryRepository(db)
	postRepository     = repositories.NewPostRepository(db)
)

func main() {
	defer func(db *sql.DB) { _ = db.Close() }(db)
	router := mux.NewRouter()

	initHandler(router, cfg)
	http.Handle("/", router)

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Project Run on " + cfg.HttpPort)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		fmt.Sprintf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		fmt.Sprintf("ctx.Done: %v", done)
	}
	cancel()
}

func initHandler(router *mux.Router, cfg *config.Config) {
	auth_handler.AuthHttpHandler(router, userRepository, cfg)
	user_handler.UserHttpHandler(router, userRepository, cfg)
	category_handler.CategoryHttpHandler(router, categoryRepository, cfg)
	post_handler.PostHttpHandler(router, postRepository, cfg)
}
