package main

import (
	"api/db"
	"api/ent"
	"api/ent/migrate"
	"api/models"
	"context"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetBookId(c *gin.Context) int {
	id := c.Param("id")
	book_id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return book_id
}

func main() {
	front_url := os.Getenv("FRONT_URL")

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{front_url}
	r.Use(cors.New(config))

	ctx := context.Background()

	client := db.OpenMariadb()
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	client.Close()

	r.GET("/books", func(c *gin.Context) {
		client := db.OpenMariadb()
		defer client.Close()
		books, err := models.GetBooks(client)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, books)
	})
	r.POST("/books", func(c *gin.Context) {
		client := db.OpenMariadb()
		defer client.Close()
		var book *ent.Book
		c.ShouldBind(&book)
		book, err := models.CreateBook(client, book)
		if err == nil {
			log.Println("book was created: ", book)
			c.JSON(200, book)
		} else {
			c.JSON(400, gin.H{"message": err.Error()})
		}
	})
	r.GET("books/:id", func(c *gin.Context) {
		client := db.OpenMariadb()
		defer client.Close()
		book_id := GetBookId(c)
		book, err := models.GetBook(client, book_id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, book)
	})
	r.PATCH("books/:id", func(c *gin.Context) {
		client := db.OpenMariadb()
		defer client.Close()
		book_id := GetBookId(c)
		var form ent.Book
		c.ShouldBind(&form)
		book, err := models.UpdateBook(client, book_id, form)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("book was updated: ", book)
		c.JSON(200, book)
	})
	r.DELETE("books/:id", func(c *gin.Context) {
		client := db.OpenMariadb()
		defer client.Close()
		book_id := GetBookId(c)
		err := models.DestroyBook(client, book_id)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.Run()
}
