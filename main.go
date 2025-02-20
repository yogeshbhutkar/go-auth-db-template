package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yogeshbhutkar/go-jwt-with-db-template/db"
	"github.com/yogeshbhutkar/go-jwt-with-db-template/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
