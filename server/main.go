package main

import (
	"bookstore/models"

	"bookstore/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //init gin - for routing

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // get all books
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run server at Port 8080
	r.Run(":8080")
}
