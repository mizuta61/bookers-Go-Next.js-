package main

import (
	"api/ent"
	"api/model"
	"api/test_responses"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
    client, err := ent.Open("mysql", "%:password@tcp(localhost:3306)/var/lib/mysql/?parseTime=True")
		if _, err = model.CreateBook(ctx, client); err != nil {
			log.Fatal(err)
		}
		book, err := model.CreateBook(ctx, client)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		defer client.Close()
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
