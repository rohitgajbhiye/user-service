package request

type RegistrationRequest struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}
