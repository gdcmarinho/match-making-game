package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"marinho/match-making-game/usecase"
)

func Expose() {
	router := gin.Default()

	router.Group("/match") 
	{
		router.POST("/find", findMatch)
	}

	router.Run(":8080")
}

func findMatch(c *gin.Context) {
	usecase.FindMatch()
	
    c.IndentedJSON(http.StatusOK, "OK")
}