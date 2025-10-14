package main

import (
	"fmt"
	"net/http"

	db "example.com/rest-api/database"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println("error->", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvents(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Values could not be parsed"})
		return
	}

	event.Id = 1
	event.UserId = 1
	err = event.Save()

	if err != nil {
		fmt.Println("error->", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event has been created", "event": event})
}
