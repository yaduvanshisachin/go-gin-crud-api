package main

import (
	"log"

	"gin-crud/database"
	"gin-crud/handler"
	"gin-crud/repository"
	"gin-crud/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepository := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService)

	// Gin
	r := gin.Default()

	// Routes
	r.GET("/users", userHandler.GetAll)
	r.GET("/users/:id", userHandler.GetByID)

	r.POST("/users", userHandler.Create)

	r.PUT("/users/:id", userHandler.Update)

	r.DELETE("/users/:id", userHandler.Delete)

	// Start server
	r.Run(":8080")
}



//go get github.com/jackc/pgx/v5