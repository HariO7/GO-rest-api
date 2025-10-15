package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/helper"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	helper.ContextErrors(err, context, "Invalid id")

	event, err := models.GetEventById(id)

	helper.ContextErrors(err, context, "An error occured")

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	helper.ContextErrors(err, context, "An error occured")

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvents(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	helper.ContextErrors(err, context, "Values could not be parsed")

	event.Id = 1
	event.UserId = 1
	err = event.Save()

	helper.ContextErrors(err, context, "An error occured")

	context.JSON(http.StatusCreated, gin.H{"message": "event has been created", "event": event})
}
