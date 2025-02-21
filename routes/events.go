package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yogeshbhutkar/go-jwt-with-db-template/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "could not parse event id"},
		)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch event"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"event": event},
	)
}

func postEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}

	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"event": event})
}

func updateEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID!"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload!"})
		return
	}

	event.ID = eventId
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}
