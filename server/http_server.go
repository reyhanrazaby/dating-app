package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanrazaby/dating-app/usecase/login"
	"github.com/reyhanrazaby/dating-app/usecase/signup"
)

func GetRoutes() http.Handler {
	router := gin.Default()
	router.Use(gin.CustomRecovery(panicHandler))
	router.Use(rateLimit())

	router.POST(login.Path, login.Handler())
	router.POST(signup.Path, signup.Handler())

	return router
}

func panicHandler(c *gin.Context, err any) {
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, gin.H{
			"messahe": "Unexpected error happened!",
		})
	}
}
