package main

import (
	"api/db"
	"api/ent/book"
	"api/model"
	"context"
	"log"
	"strconv"

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
		ctx := context.Background()
    client := db.Main()
		defer client.Close()
		books, err := client.Book.     
    Query().                   
    All(ctx)
		if err != nil {
			log.Fatal(err)
	  }
    c.JSON(200, books)
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
	
	r.GET("books/:id", func(c *gin.Context)  {
		id := c.Param("id")
		var book_id int
		book_id, _ = strconv.Atoi(id) 
		ctx := context.Background()
    client := db.Main()
		defer client.Close()
		book, err := client.Book.     
    Query().  
		Where(book.ID(book_id)).                 
    Only(ctx)
		if err != nil {
			log.Fatal(err)
	  }
		c.JSON(200, book)
	})
	r.PATCH("books/:id", func(c *gin.Context)  {
		id := c.Param("id") 
		var book_id int
		book_id, _ = strconv.Atoi(id) 
		ctx := context.Background()
    client := db.Main()
		defer client.Close()
		book, err := model.UpdateBook(ctx, client, c, book_id)
		if err != nil {
			log.Fatalf("failed opening connection to mysql:db %v", err)
		}
		c.JSON(200, book)
	})
	r.DELETE("books/:id", func(c *gin.Context)  {
		id := c.Param("id") 
		var book_id int
		book_id, _ = strconv.Atoi(id)
		ctx := context.Background()
    client := db.Main()
		defer client.Close()
		err := client.Book.
    DeleteOneID(book_id).
    Exec(ctx)
		if err != nil {
			log.Fatal(err)
	  }
	})

	r.Run()
}
