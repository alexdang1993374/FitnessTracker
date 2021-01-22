package routes

import (
	"net/http"

	"fitness-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/api", welcome)
	router.GET("/api/users", controllers.GetAllUsers)
	router.POST("/api/users", controllers.CreateUser)
	router.GET("/api/users/:userId", controllers.GetSingleUser)
	router.PUT("/api/users/:userId", controllers.EditUser)
	router.DELETE("/api/users/:userId", controllers.DeleteUser)

	router.GET("/api/exercises", controllers.GetAllExercises)
	router.POST("/api/exercises", controllers.CreateExercise)
	router.GET("/api/exercises/:exercisesId", controllers.GetSingleExercise)
	router.PUT("/api/exercises/:exercisesId", controllers.EditExercise)
	router.DELETE("/api/exercises/:exercisesId", controllers.DeleteExercise)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}
