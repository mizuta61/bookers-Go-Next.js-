package main

import (
	"api/test_responses"
	"fmt"

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
		id := c.Query("id")
		title := c.PostForm("title")
		body := c.PostForm("body")

		fmt.Printf("id: %s; title: %s; body: %s", id, title, body)
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
