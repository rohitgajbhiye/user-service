package main

import (
	"fmt"
	"net/http"
	"os"
	"user-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// store.NewRedisStore()
	r := gin.Default()
	router.InitRoute(r)
	port := os.Getenv("port")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("Error initializing router", err.Error())
	}
}
