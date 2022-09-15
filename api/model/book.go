package model

import (
	"api/ent"
	"api/ent/book"
	"context"
	"fmt"
	"log"
	"time"
)

type Book struct {
	Id        int        `json:"id" form:"id"`
	Title     string     `json:"title" form:"title"`
	Body      string     `json:"body" form:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func CreateBook(ctx context.Context, client *ent.Client) (*ent.Book, error){
	book, err := client.Book.
	Create().
	SetTitle(book.FieldTitle).
	SetBody(book.FieldBody).
	Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("book was created: ", book)
	return book, err
}