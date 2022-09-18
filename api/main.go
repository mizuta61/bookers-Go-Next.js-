package main

import (
	"api/models"
	"context"

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
		models.CreateBook(ctx, c)	
  })
	r.GET("books/:id", func(c *gin.Context)  { 
		models.GetBook(ctx, c)
	})
	r.PATCH("books/:id", func(c *gin.Context)  {   
		models.UpdateBook(ctx, c)
	})
	r.DELETE("books/:id", func(c *gin.Context)  {
		models.DestroyBook(ctx, c)
	})

	r.Run()
}
