package main

import (
	"net/http"

	"github.com/reyhanrazaby/dating-app/usecase/login"
	"github.com/reyhanrazaby/dating-app/usecase/signup"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	server := &http.Server{
		Addr:    ":4545",
		Handler: setUpRoutes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}

func setUpRoutes() http.Handler {
	router := gin.Default()
	router.Use(gin.CustomRecovery(panicHandler))

	router.POST(login.Path, login.Handler())
	router.POST(signup.Path, signup.Handler())

	return router
}

type SignUpResponse struct {
	Message string `json:"message"`
}

func panicHandler(c *gin.Context, err any) {
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, gin.H{
			"errors": []string{"Unexpected error happened!"},
		})
	}
}
