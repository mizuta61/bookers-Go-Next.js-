package models

import (
	"api/ent"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id        int        `json:"id" form:"id"`
	Title     string     `json:"title" form:"title"`
	Body      string     `json:"body" form:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func CreateBook(ctx context.Context, client *ent.Client, c *gin.Context) (*ent.Book, error){
	var form Book
  c.ShouldBind(&form)
	book, err := client.Book.
	Create().
	SetTitle(form.Title).
	SetBody(form.Body).
	Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("book was created: ", book)
	return book, err
}

func UpdateBook(ctx context.Context, client *ent.Client, c *gin.Context) (*ent.Book, error){
	id := c.Param("id") 
	var book_id int
	book_id, _ = strconv.Atoi(id) 
	var form Book
  c.ShouldBind(&form)
	book, err := client.Book.
	UpdateOneID(book_id). 
	SetTitle(form.Title).
	SetBody(form.Body).
	Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("book was updated: ", book)
	return book, err
}

func DestroyBook(ctx context.Context, client *ent.Client, c *gin.Context) {
	id := c.Param("id") 
	var book_id int
	book_id, _ = strconv.Atoi(id)
	err := client.Book.
	DeleteOneID(book_id).
	Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}