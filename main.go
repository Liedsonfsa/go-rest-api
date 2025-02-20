package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.DB, _ = db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Event created!", "event": event})
}