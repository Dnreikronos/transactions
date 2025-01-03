package routes

import (
	"github.com/Dnreikronos/transactions/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/createTransactions", handlers.CreateTransaction)
}
