package mapper

import (
	"user-service/controller/request"
	"user-service/model"
)

func RegistrationToUserMapper(request request.RegistrationRequest) model.User {
	return model.User{
		FirstName: request.FirstName,
		Email:     request.Email,
	}
}
