package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(constext *gin.Context) {
	events := models.GetAllEvents()
	constext.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvents(context *gin.Context) {
	var event models.Events

	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Values could not be parsed"})
		return
	}

	event.Id = 1
	event.UserId = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event has been created", "event": event})
}
