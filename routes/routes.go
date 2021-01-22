package routes

import (
	"net/http"

	"fitness-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/api", welcome)
	router.GET("/api/fitness", controllers.GetAllFitness)
	router.POST("/api/fitness", controllers.CreateFitness)
	router.GET("/api/fitness/:userId", controllers.GetSingleFitness)
	router.PUT("/api/fitness/:userId", controllers.EditFitness)
	router.DELETE("/api/fitness/:userId", controllers.DeleteFitness)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}
