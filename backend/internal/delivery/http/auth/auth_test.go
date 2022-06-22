package auth_http_handler_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"backend/internal/config"
	auth_http_handler "backend/internal/delivery/http/auth"
	user_http_handler "backend/internal/delivery/http/user"
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

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	createdAt string `json:"created_at"`
	updatedAt string `json:"updated_at"`
}

type ResponseRegister struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

var (
	cfg = config.InitConfig()
)

var _ = Describe("Auth", func() {
	When("Login", func() {

		It("Should return data user when logged in", func() {
			// Given
			router := mux.NewRouter()
			db, err := sql.Open("sqlite3", "../../../../halobelajar.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			bodyReader := strings.NewReader(`{"email":"ruang2@.ac.id","password":"ruang2"}`)

			userRepository := repository.NewUserRepository(db)

			route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(200))
		})

		It("Should return error when credentials is invalid", func() {
			router := mux.NewRouter()
			db, err := sql.Open("sqlite3", "../../../../halobelajar.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			bodyReader := strings.NewReader(`{"email":"ruang2@.ac.id","password":"ppasswordsalah"}`)

			userRepository := repository.NewUserRepository(db)

			route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			route.ServeHTTP(w, r)
			//Expect(w.Code).To(Equal(http.StatusUnauthorized))
			Expect(w.Body.String()).To(Equal(`{"code":500,"message":"Error","errors":["incorrect password. please enter a password that matches your registered email"]}`))
		})
	})
	When("Register", func() {
		It("Should return data user when registered", func() {
			router := mux.NewRouter()
			db, err := sql.Open("sqlite3", "../../../../halobelajar.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			bodyReader := strings.NewReader(`{"name":"ruang6","email":"ruang6@.ac.id","password":"ruang66"}`)

			userRepository := repository.NewUserRepository(db)

			route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/api/auth/register", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			route.ServeHTTP(w, r)
			var response ResponseRegister
			err = json.Unmarshal([]byte(w.Body.String()), &response)

			Expect(w.Code).To(Equal(201))

			id := response.Data.ID

			route2 := user_http_handler.UserHttpHandler(router, userRepository, cfg)
			r, err = http.NewRequest("DELETE", "/api/user/"+id, nil)
			route2.ServeHTTP(w, r)
		})
	})

})

// func TestLogin(t *testing.T) {
// 	router := mux.NewRouter()
// 	db, err := sql.Open("sqlite3", "../../../../halobelajar.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	bodyReader := strings.NewReader(`{"email":"ruang2@.ac.id","password":"ruang2"}`)

// 	userRepository := repository.NewUserRepository(db)
// 	request, _ := http.NewRequest("POST", "/api/auth/login", bodyReader)
// 	response := httptest.NewRecorder()

// 	auth_http_handler.AuthHttpHandler(router, userRepository, cfg).ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code)
// 	http.Handle("/", router)

// }
