package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.DB, _ = db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(id)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event."})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Event created!", "event": event})
}