package models

import (
	"api/db"
	"api/ent"
	"api/ent/book"
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

func DestroyBook(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
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

func GetBook(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
	id := c.Param("id")
	var book_id int
	book_id, _ = strconv.Atoi(id) 
	book, err := client.Book.     
	Query().  
	Where(book.ID(book_id)).                 
	Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, book)
}

func BookLists(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
	books, err := client.Book.     
	Query().                   
	All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, books)
}