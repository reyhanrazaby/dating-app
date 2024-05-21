package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanrazaby/dating-app/errors"
	"github.com/reyhanrazaby/dating-app/util"
)

var Path = "/logout"
var service SignUpService = NewService()

func Handler() func(*gin.Context) {
	return func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			util.ErrorJson(c, http.StatusBadRequest, err)
			return
		}

		err = service.SignUp(req)
		if err != nil {
			var errorCode int
			switch err.(type) {
			case errors.SignUpError:
				errorCode = http.StatusBadRequest
			default:
				errorCode = http.StatusInternalServerError
			}
			util.ErrorJson(c, errorCode, err)
			return
		}

		json := response{
			Message: "User created successfully",
		}
		c.JSON(http.StatusOK, json)
	}
}

type request struct {
	FullName string `json:"full_name" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type response struct {
	Message string `json:"message"`
}
