package auth_http_handler

import (
	"database/sql"
	"strings"

	"backend/internal/config"
	"backend/internal/repository"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TokenStr struct {
	Token string `json:"token"`
}

type loginResp struct {
	Code    int      `json:"code"`
	Data    TokenStr `json:"data"`
	Message string   `json:"message"`
}

var (
	cfg = config.InitConfig()
)
var _ = Describe("Auth", func() {
	It("Should return data user when logged in", func() {
		// Given
		router := mux.NewRouter()
		db, err := sql.Open("sqlite3", "../../../halobelajar.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		bodReader := strings.NewReader(`{"email":"ruang2@.ac.id","password":"ruang2"}`)
		userRepository := repository.NewUserRepository(db)
		main.initHandler(router, cfg)
		// When
		// Then
	})
})
