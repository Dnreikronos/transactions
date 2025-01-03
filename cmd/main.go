package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/transactions/configs"
	"github.com/Dnreikronos/transactions/db/connection"
	"github.com/Dnreikronos/transactions/db/migrations"
	"github.com/Dnreikronos/transactions/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	db := connection.OpenConnection()
	migrations.RunMigrations(db)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.RegisterRoutes(r)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
