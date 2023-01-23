package Auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pascalallen/baetyl/src/Adapter/Http/Responder/JSend"
	"github.com/pascalallen/baetyl/src/Domain/Auth/User"
	"net/http"
)

func InternalServerErrorResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusInternalServerError,
		JSend.ErrorResponse[string]{
			Status:  "error",
			Message: error.Error(),
		},
	)

	return
}

func UnprocessableEntityResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusUnprocessableEntity,
		JSend.ErrorResponse[string]{
			Status:  "error",
			Message: error.Error(),
		},
	)

	return
}

func BadRequestResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusBadRequest,
		JSend.FailResponse[string]{
			Status: "fail",
			Data:   error.Error(),
		},
	)

	return
}

func CreatedResponse(c *gin.Context, user *User.User) {
	bytes, _ := json.Marshal(user)

	c.JSON(
		http.StatusCreated,
		JSend.SuccessResponse[string]{
			Status: "success",
			Data:   string(bytes),
		},
	)

	return
}
