package Auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	RegisterUserResponder "github.com/pascalallen/Baetyl/src/Adapter/Http/Responder/Api/V1/Auth"
	GormUserRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/User"
	UserValidations "github.com/pascalallen/Baetyl/src/Adapter/Validation/Auth/User"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/PasswordHash"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
)

func Handle(c *gin.Context) {
	var request UserValidations.RegisterUserRules

	if err := c.ShouldBind(&request); err != nil {
		RegisterUserResponder.BadRequestResponse(c, err)

		return
	}

	userRepository := GormUserRepository.GormUserRepository{}
	if user, _ := userRepository.GetByEmailAddress(request.EmailAddress); user != nil {
		RegisterUserResponder.UnprocessableEntityResponse(c, fmt.Errorf("user already exists with email address: %s", request.EmailAddress))

		return
	}

	id := ulid.Make()
	user := User.Register(id, request.FirstName, request.LastName, request.EmailAddress)
	passwordHash := PasswordHash.Create(request.Password)
	user.SetPasswordHash(passwordHash)

	if err := userRepository.Add(user); err != nil {
		RegisterUserResponder.InternalServerErrorResponse(c, err)

		return
	}

	RegisterUserResponder.CreatedResponse(c, user)

	return
}
