package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/reyhanrazaby/dating-app/datasource/inmemory"
	"github.com/reyhanrazaby/dating-app/errors"
	"github.com/reyhanrazaby/dating-app/util"
)

var Path = "/login"
var service LoginService = NewService(inmemory.GetInstance())

func Handler() func(*gin.Context) {
	return func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			util.ErrorJson(c, http.StatusBadRequest, err)
			return
		}

		userProfile, err := service.Login(req.Email, req.Password)
		if err != nil {
			var errorCode int
			switch err.(type) {
			case errors.AuthError:
				errorCode = http.StatusBadRequest
			default:
				errorCode = http.StatusInternalServerError
			}
			util.ErrorJson(c, errorCode, err)
			return
		}

		json := response{
			UserId:  userProfile.Id,
			Message: "Login successfully",
		}
		c.JSON(http.StatusOK, json)
	}
}

type request struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type response struct {
	UserId  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}
