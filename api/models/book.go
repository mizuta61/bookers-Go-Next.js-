package models

import (
	"api/db"
	"api/ent"
	"api/ent/book"
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookId(c *gin.Context) int {
	id := c.Param("id")
	book_id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return book_id
}

func CreateBook(client *ent.Client, book *ent.Book) (*ent.Book, error) {
	ctx := context.Background()
	book, err := client.Book.
		Create().
		SetTitle(book.Title).
		SetBody(book.Body).
		Save(ctx)
	return book, err
}

func UpdateBook(ctx context.Context, c *gin.Context) (*ent.Book, error) {
	client := db.OpenMariadb()
	defer client.Close()
	book_id := GetBookId(c)
	var form ent.Book
	c.ShouldBind(&form)
	book, err := client.Book.
		UpdateOneID(book_id).
		SetTitle(form.Title).
		SetBody(form.Body).
		Save(ctx)
	return book, err
}

func DestroyBook(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
	book_id := GetBookId(c)
	err := client.Book.
		DeleteOneID(book_id).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBook(client *ent.Client, book_id int) (*ent.Book, error) {
	ctx := context.Background()
	book, err := client.Book.
		Query().
		Where(book.ID(book_id)).
		Only(ctx)
	return book, err
}

func GetBooks(ctx context.Context, c *gin.Context) ([]*ent.Book, error) {
	client := db.OpenMariadb()
	defer client.Close()
	books, err := client.Book.
		Query().
		All(ctx)
	return books, err
}
