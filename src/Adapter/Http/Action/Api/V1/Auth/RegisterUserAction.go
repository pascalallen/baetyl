package Auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	RegisterUserResponder "github.com/pascalallen/baetyl/src/Adapter/Http/Responder/Api/V1/Auth"
	GormUserRepository "github.com/pascalallen/baetyl/src/Adapter/Repository/Auth/User"
	UserValidations "github.com/pascalallen/baetyl/src/Adapter/Validation/Auth/User"
	"github.com/pascalallen/baetyl/src/Domain/Auth/PasswordHash"
	"github.com/pascalallen/baetyl/src/Domain/Auth/User"
)

func Handle(c *gin.Context) {
	var request UserValidations.RegisterUserRules

	if err := c.ShouldBind(&request); err != nil {
		errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
		RegisterUserResponder.BadRequestResponse(c, errors.New(errorMessage))

		return
	}

	var userRepository User.UserRepository = GormUserRepository.GormUserRepository{}

	if user, err := userRepository.GetByEmailAddress(request.EmailAddress); user != nil || err != nil {
		errorMessage := fmt.Sprint("Something went wrong. If you already have an account, please log in.")
		RegisterUserResponder.UnprocessableEntityResponse(c, errors.New(errorMessage))

		return
	}

	id := ulid.Make()
	user := User.Register(id, request.FirstName, request.LastName, request.EmailAddress)
	passwordHash := PasswordHash.Create(request.Password)
	user.SetPasswordHash(passwordHash)

	if err := userRepository.Add(user); err != nil {
		errorMessage := fmt.Sprintf("Error persisting user: %s", err.Error())
		RegisterUserResponder.InternalServerErrorResponse(c, errors.New(errorMessage))

		return
	}

	RegisterUserResponder.CreatedResponse(c, user)

	return
}
