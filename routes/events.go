package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("error->", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid id"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		fmt.Println("error->", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
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
