package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Register health check route.
	server.GET("/", healthCheck)

	// Register events routes.
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", postEvents)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
