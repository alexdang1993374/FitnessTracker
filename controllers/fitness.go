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

//User Table created
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Exercise Table created
type Exercise struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//CreateUserTable makes the user table
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

//CreateExerciseTable makes the exercise table
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

//GetAllUsers in request
func GetAllUsers(c *gin.Context) {
	var users []User
	err := dbConnect.Model(&users).Select()
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
		"data":    users,
	})
	return
}

//CreateUser request
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

//GetSingleUser request
func GetSingleUser(c *gin.Context) {
	userID := c.Param("userId")
	user := &User{ID: userID}
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

//EditUser request
func EditUser(c *gin.Context) {
	userID := c.Param("userId")
	var user User
	c.BindJSON(&user)
	username := user.Username
	_, err := dbConnect.Model(&User{}).Set("username = ?", username).Where("id = ?", userID).Update()
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

//DeleteUser request
func DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	user := &User{ID: userID}
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

//Requests for Exercise table

//GetAllExercises request
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

//CreateExercise request
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
		Date:        time.Now(),
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

//GetSingleExercise request
func GetSingleExercise(c *gin.Context) {
	exerciseID := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseID}
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

//EditExercise request
func EditExercise(c *gin.Context) {
	exerciseID := c.Param("exerciseId")
	var exercise Exercise
	c.BindJSON(&exercise)
	username := exercise.Username
	description := exercise.Description
	duration := exercise.Duration
	date := exercise.Date
	_, err := dbConnect.Model(&Exercise{}).Set("username = ?", username).Where("id = ?", exerciseID).Update()
	_, err = dbConnect.Model(&Exercise{}).Set("description = ?", description).Where("id = ?", exerciseID).Update()
	_, err = dbConnect.Model(&Exercise{}).Set("duration = ?", duration).Where("id = ?", exerciseID).Update()
	_, err = dbConnect.Model(&Exercise{}).Set("date = ?", date).Where("id = ?", exerciseID).Update()
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

//DeleteExercise request
func DeleteExercise(c *gin.Context) {
	exerciseID := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseID}
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

//InitiateDB initiates the database
func InitiateDB(db *pg.DB) {
	dbConnect = db
}
