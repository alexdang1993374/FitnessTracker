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

type Fitness struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed int       `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create User Table
func CreateFitnessTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Fitness{}, opts)
	if createError != nil {
		log.Printf("Error while creating fitness table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Fitness table created")
	return nil
}

func GetAllFitness(c *gin.Context) {
	var fitness []Fitness
	err := dbConnect.Model(&fitness).Select()
	if err != nil {
		log.Printf("Error while getting all fitness, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Fitness",
		"data":    fitness,
	})
	return
}

func CreateFitness(c *gin.Context) {
	var fitness Fitness
	c.BindJSON(&fitness)
	title := fitness.Title
	body := fitness.Body
	completed := fitness.Completed
	id := guuid.New().String()
	insertError := dbConnect.Insert(&Fitness{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new fitness into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Fitness created Successfully",
	})
	return
}
func GetSingleFitness(c *gin.Context) {
	fitnessId := c.Param("userId")
	fitness := &Fitness{ID: fitnessId}
	err := dbConnect.Select(fitness)
	if err != nil {
		log.Printf("Error while getting a single fitness, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Fitness not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Fitness",
		"data":    fitness,
	})
	return
}
func EditFitness(c *gin.Context) {
	fitnessId := c.Param("userId")
	var fitness Fitness
	c.BindJSON(&fitness)
	completed := fitness.Completed
	_, err := dbConnect.Model(&Fitness{}).Set("completed = ?", completed).Where("id = ?", fitnessId).Update()
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
		"message": "Fitness Edited Successfully",
	})
	return
}
func DeleteFitness(c *gin.Context) {
	fitnessId := c.Param("userId")
	fitness := &Fitness{ID: fitnessId}
	err := dbConnect.Delete(fitness)
	if err != nil {
		log.Printf("Error while deleting a single fitness, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Fitness deleted successfully",
	})
	return
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
