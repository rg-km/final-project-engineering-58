package request

type CreateUserPayload struct {
	Name	string `json:"name,omitempty"`
	Email	string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role	string `json:"role,omitempty"`
}