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

func CreateBook(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
	var form ent.Book
	c.ShouldBind(&form)
	book, err := client.Book.
		Create().
		SetTitle(form.Title).
		SetBody(form.Body).
		Save(ctx)
	if err == nil {
		log.Println("book was created: ", book)
		c.JSON(200, book)
	} else {
		c.JSON(400, gin.H{"message": err.Error()})
	}

}

func UpdateBook(ctx context.Context, c *gin.Context) {
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
	if err != nil {
		log.Fatal(err)
	}
	log.Println("book was updated: ", book)
	c.JSON(200, book)
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

func GetBook(ctx context.Context, c *gin.Context) {
	client := db.OpenMariadb()
	defer client.Close()
	book_id := GetBookId(c)
	book, err := client.Book.
		Query().
		Where(book.ID(book_id)).
		Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, book)
}

func GetBooks(ctx context.Context, c *gin.Context) {
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
