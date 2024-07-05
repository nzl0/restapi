package main

import (
	"github.com/gin-gonic/gin"
	"library2/internal/domain/model"
	"library2/internal/domain/service"
	"library2/internal/handler"
	"library2/pkg/db"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//database connection created

	dB, err := db.Connection()
	if err != nil {
		panic(err)
	}

	//database schemas provided
	err = dB.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatalf("Auto migrate book err: %v", err)
	}
	//to database data added
	err = dB.Seed()
	if err != nil {
		log.Fatalf("Seed book err: %v", err)
	}
	bookRepo := db.NewBookRepository(dB.GetDb())
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/books", bookHandler.GetAll)
	v1.GET("/books/filter", bookHandler.GetAllWithFilter)
	v1.GET("/books/id/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
