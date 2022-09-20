package main

import (
	"api/db"
	"api/ent/migrate"
	"api/models"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	front_url := os.Getenv("FRONT_URL")

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{front_url}
	r.Use(cors.New(config))

	ctx := context.Background()

	client := db.OpenMariadb()
	defer client.Close()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
