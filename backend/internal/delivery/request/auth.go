package request

type CreateAuthPayload struct {
	Email   	string `json:"email"`
	Password 	string `json:"password"`
}