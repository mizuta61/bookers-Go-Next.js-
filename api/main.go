package main

import (
	"api/db"
	"api/models"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func main()  {
	r := gin.Default()
	config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:8000"}
  r.Use(cors.New(config))

	ctx := context.Background()

	r.GET("/books", func(c *gin.Context)  {
    models.BookLists(ctx, c)
	})
	r.POST("/books", func(c *gin.Context) {
    client := db.OpenMariadb()
		defer client.Close()
		book, err := models.CreateBook(ctx, client, c)
		if err != nil {
			log.Fatalf("failed opening connection to mysql:db %v", err)
		}
		c.JSON(200, book)
  })
	
	r.GET("books/:id", func(c *gin.Context)  { 
		models.GetBook(ctx, c)
	})
	r.PATCH("books/:id", func(c *gin.Context)  {
    client := db.OpenMariadb()
		defer client.Close()
		book, err := models.UpdateBook(ctx, client, c)
		if err != nil {
			log.Fatalf("failed opening connection to mysql:db %v", err)
		}
		c.JSON(200, book)
	})
	r.DELETE("books/:id", func(c *gin.Context)  {
		models.DestroyBook(ctx, c)
	})

	r.Run()
}
