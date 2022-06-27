package main

import (
	"fmt"
	"heroku-deployment/config"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")

	filename := "config_dev.json"

	if port == "" {
		log.Fatalf("provide PORT must be set")
	}

	if mode == "release" {
		filename = "config_prod.json"
		gin.SetMode(gin.ReleaseMode)
	}

	err := config.SetupConfig(filename)

	if err != nil {
		log.Fatalf("error in setup config file")
	}

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
