package auth_http_handler_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

type ResponseRegister struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var (
	cfg = config.InitConfig()
)

var responseUseforDelete ResponseRegister
var responseRegister ResponseRegister

var _ = Describe("Auth", func() {
	Describe("Register", func() {
		When("Email and Password are valid", func() {
			It("Should return data user when registered", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"name":"percobaan","email":"percobaan@.ac.id","password":"percobaan"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/register", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				err = json.Unmarshal([]byte(w.Body.String()), &responseRegister)

				Expect(responseRegister.Data.Email).To(Equal("percobaan@.ac.id"))
				Expect(responseRegister.Data.Name).To(Equal("percobaan"))
			})
			It("Should return code 200 and Message Success", func() {
				Expect(responseRegister.Code).To(Equal(201))
				Expect(responseRegister.Message).To(Equal("Success"))
			})
		})

		When("Email has already registered", func() {
			It("Should return Internal Server Error", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"name":"percobaan","email":"percobaan@.ac.id","password":"percobaan"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/register", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				var response ResponseError
				err = json.Unmarshal([]byte(w.Body.String()), &response)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})

			It("Should return message Email already exist", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"name":"percobaan","email":"percobaan@.ac.id","password":"percobaan"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/register", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				var response ResponseError
				err = json.Unmarshal([]byte(w.Body.String()), &response)
				Expect(response.Errors[0]).To(Equal("email already exist"))
			})
		})
	})

	Describe("Login", func() {
		When("Email and Password are valid", func() {
			It("Should return data user", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"email":"percobaan@.ac.id","password":"percobaan"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)
				err = json.Unmarshal([]byte(w.Body.String()), &responseUseforDelete)

				Expect(responseUseforDelete.Data.Email).To(Equal("percobaan@.ac.id"))
			})
			It("Should return code 200 and success message", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"email":"percobaan@.ac.id","password":"percobaan"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				Expect(responseUseforDelete.Code).To(Equal(200))
				Expect(responseUseforDelete.Message).To(Equal("Success"))
			})
		})
		When("Email valid but Password are invalid", func() {
			It("Should return Internal Server Error and error message", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"email":"percobaan@.ac.id","password":"passwordsalah"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				var response ResponseError
				err = json.Unmarshal([]byte(w.Body.String()), &response)

				Expect(response.Code).To(Equal(http.StatusInternalServerError))
				Expect(response.Errors[0]).To(Equal("incorrect password. please enter a password that matches your registered email"))
			})
		})
		When("Both Email and Password are invalid", func() {
			It("Should return error when email is incorrect and password incorrect", func() {
				router := mux.NewRouter()
				db, err := sql.Open("sqlite3", "./database_test/test.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()

				bodyReader := strings.NewReader(`{"email":"emailsalah@.ac.id","password":"passwordsalah"}`)

				userRepository := repository.NewUserRepository(db)

				route := auth_http_handler.AuthHttpHandler(router, userRepository, cfg)
				w := httptest.NewRecorder()
				r, err := http.NewRequest("POST", "/api/auth/login", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				route.ServeHTTP(w, r)

				var response ResponseError
				err = json.Unmarshal([]byte(w.Body.String()), &response)
				if err != nil {
					log.Fatal(err)
				}

				Expect(response.Code).To(Equal(http.StatusInternalServerError))
				Expect(response.Errors[0]).To(Equal("invalid email"))
				Expect(response.Message).To(Equal("Error"))

				//DELETE USER DATA FOR NEXT TEST
				id := responseUseforDelete.Data.ID
				fmt.Println("DELETED ID :", id)

				routerX := mux.NewRouter()
				dbX, errX := sql.Open("sqlite3", "./database_test/test.db")
				if errX != nil {
					panic(errX)
				}
				defer dbX.Close()

				userRepositoryX := repository.NewUserRepository(dbX)

				routeX := user_http_handler.UserHttpHandler(routerX, userRepositoryX, cfg)
				wX := httptest.NewRecorder()

				rX, errX := http.NewRequest("DELETE", "/api/user/"+id, nil)
				if errX != nil {
					log.Fatal(errX)
				}
				routeX.ServeHTTP(wX, rX)
			})
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
// 	assert.Equal(t, 201, response.Code)
// 	http.Handle("/", router)

// }
