package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalf("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"health": "UP",
		})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error in setup :", err.Error())
	}

	fmt.Printf("Server runnig on port %+v :", port)

}
