package Auth

import (
	"github.com/gin-gonic/gin"
	Responder "github.com/pascalallen/Baetyl/src/Adapter/Http/Responder/Api/V1/Auth"
	UserValidations "github.com/pascalallen/Baetyl/src/Adapter/Validation/Auth/User"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/PasswordHash"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
)

func Handle(c *gin.Context) {
	var registerUserRules UserValidations.RegisterUserRules

	err := c.ShouldBind(&registerUserRules)
	if err != nil {
		Responder.BadRequestResponse(c, err)

		return
	}

	user := User.Register(registerUserRules.FirstName, registerUserRules.LastName, registerUserRules.EmailAddress)
	passwordHash := PasswordHash.Create(registerUserRules.Password)
	user.SetPasswordHash(passwordHash)

	Responder.CreatedResponse(c, user)

	return
}
