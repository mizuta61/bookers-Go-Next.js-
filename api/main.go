package main

import (
	"api/db"
	"api/model"
	"api/test_responses"
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

	
	r.GET("/books", func(c *gin.Context)  {
		c.JSON(200, test_responses.TestBooks())
	})
	r.POST("/books", func(c *gin.Context) {
		ctx := context.Background()
    client := db.Main()
		defer client.Close()
		book, err := model.CreateBook(ctx, client, c)
		if err != nil {
			log.Fatalf("failed opening connection to mysql:db %v", err)
		}
		c.JSON(200, book)
  })
	
	// r.GET("books/:id", func(c *gin.Context)  {
	// 	id := c.Param("id") 
	// })
	// r.PATCH("books/:id", func(c *gin.Context)  {
	// 	id := c.Param("id") 
	// })
	// r.DELETE("books/:id", func(c *gin.Context)  {
	// 	id := c.Param("id") 
	// })

	r.Run()
}
