package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pactice_02/model"
)

var Books []model.Book

func main() {
	log.Println("Start Server ...")

	router := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

}
