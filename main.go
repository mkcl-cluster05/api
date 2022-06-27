package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"health": "UP",
		})
	})

	return router
}

func main() {

	router := setupRouter()
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")

	if err := router.Run(port); err != nil {
		log.Fatal("Error in setup :", err.Error())
	}

}
