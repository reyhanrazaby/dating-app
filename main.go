package main

import (
	"net/http"

	"github.com/reyhanrazaby/dating-app/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":4545",
		Handler: server.GetRoutes(),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}
