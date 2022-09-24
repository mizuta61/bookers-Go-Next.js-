package main

import (
	"api/db"
	"api/ent/migrate"
	"api/models"
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
		models.GetBooks(ctx, c)
	})
	r.POST("/books", func(c *gin.Context) {
		models.CreateBook(ctx, c)
	})
	r.GET("books/:id", func(c *gin.Context) {
		models.GetBook(ctx, c)
	})
	r.PATCH("books/:id", func(c *gin.Context) {
		models.UpdateBook(ctx, c)
	})
	r.DELETE("books/:id", func(c *gin.Context) {
		models.DestroyBook(ctx, c)
	})

	r.Run()
}
