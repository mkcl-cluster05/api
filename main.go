package main

import (
	"log"
	"net/http"

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

	if err := router.Run(":3000"); err != nil {
		log.Fatal("Error in setup :", err.Error())
	}

}
