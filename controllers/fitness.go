package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Exercise struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&User{}, opts)
	if createError != nil {
		log.Printf("Error while creating user table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("User table created")
	return nil
}

func CreateExerciseTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Exercise{}, opts)
	if createError != nil {
		log.Printf("Error while creating exercise table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Exercise table created")
	return nil
}

func GetAllUsers(c *gin.Context) {
	var user []User
	err := dbConnect.Model(&user).Select()
	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    user,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	username := user.Username
	id := guuid.New().String()
	insertError := dbConnect.Insert(&User{
		ID:        id,
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}
func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID: userId}
	err := dbConnect.Select(user)
	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})
	return
}
func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	var user User
	c.BindJSON(&user)
	username := user.Username
	_, err := dbConnect.Model(&User{}).Set("username = ?", username).Where("id = ?", userId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Edited Successfully",
	})
	return
}
func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID: userId}
	err := dbConnect.Delete(user)
	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}

//exercise crud

func GetAllExercises(c *gin.Context) {
	var exercise []Exercise
	err := dbConnect.Model(&exercise).Select()
	if err != nil {
		log.Printf("Error while getting all exercises, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Exercises",
		"data":    exercise,
	})
	return
}

func CreateExercise(c *gin.Context) {
	var exercise Exercise
	c.BindJSON(&exercise)
	username := exercise.Username
	description := exercise.Description
	duration := exercise.Duration
	id := guuid.New().String()
	insertError := dbConnect.Insert(&Exercise{
		ID:          id,
		Username:    username,
		Description: description,
		Duration:    duration,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new exercise into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Exercise created Successfully",
	})
	return
}
func GetSingleExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}
	err := dbConnect.Select(exercise)
	if err != nil {
		log.Printf("Error while getting a single exercise, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Exercise not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Exercise",
		"data":    exercise,
	})
	return
}
func EditExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
	var exercise Exercise
	c.BindJSON(&exercise)
	username := exercise.Username
	description := exercise.Description
	duration := exercise.Duration
	_, err := dbConnect.Model(&User{}).Set("username = ?", username).Where("id = ?", exerciseId).Update()
	_, err = dbConnect.Model(&User{}).Set("description = ?", description).Where("id = ?", exerciseId).Update()
	_, err = dbConnect.Model(&User{}).Set("duration = ?", duration).Where("id = ?", exerciseId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Exercise Edited Successfully",
	})
	return
}
func DeleteExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}
	err := dbConnect.Delete(exercise)
	if err != nil {
		log.Printf("Error while deleting a single exercise, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Exercise deleted successfully",
	})
	return
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
