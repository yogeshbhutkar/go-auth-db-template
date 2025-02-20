package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Health check successful!"})
}
