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
	if helper.ContextErrors(err, context, http.StatusBadRequest, "Invalid id") {
		return
	}

	event, err := models.GetEventById(id)
	if helper.ContextErrors(err, context, http.StatusNotFound, "Event not found") {
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if helper.ContextErrors(err, context, http.StatusBadRequest, "Invalid request") {
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvents(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Values could not be parsed") {
		return
	}

	event.Id = 1
	err = event.Save()
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "An error occured") {
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event has been created", "event": event})
}

func updateEventById(context *gin.Context) {

	//checking if id exists
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Invalid id") {
		return
	}

	//check if event is present
	_, err = models.GetEventById(id)
	if helper.ContextErrors(err, context, http.StatusNotFound, "Event not found") {
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Could not parse values") {
		return
	}

	updatedEvent.Id = id
	err = updatedEvent.Update()
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Updation failed") {
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updation successfull"})

}

func deleteEventById(context *gin.Context) {
	//checking if id exists
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Invalid id") {
		return
	}

	//check if event is present
	_, err = models.GetEventById(id)
	if helper.ContextErrors(err, context, http.StatusNotFound, "Event not found") {
		return
	}

	var deletedEvent models.Event
	deletedEvent.Id = id
	err = deletedEvent.Delete()
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Updation failed") {
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deletion successfull"})
}
