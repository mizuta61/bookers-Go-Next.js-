package models

import (
	"api/ent"
	"api/ent/book"
	"context"
)

func CreateBook(client *ent.Client, book *ent.Book) (*ent.Book, error) {
	ctx := context.Background()
	book, err := client.Book.
		Create().
		SetTitle(book.Title).
		SetBody(book.Body).
		Save(ctx)
	return book, err
}

func UpdateBook(client *ent.Client, book_id int, form ent.Book) (*ent.Book, error) {
	ctx := context.Background()
	book, err := client.Book.
		UpdateOneID(book_id).
		SetTitle(form.Title).
		SetBody(form.Body).
		Save(ctx)
	return book, err
}

func DestroyBook(client *ent.Client, book_id int) error {
	ctx := context.Background()
	err := client.Book.
		DeleteOneID(book_id).
		Exec(ctx)
	return err
}

func GetBook(client *ent.Client, book_id int) (*ent.Book, error) {
	ctx := context.Background()
	book, err := client.Book.
		Query().
		Where(book.ID(book_id)).
		Only(ctx)
	return book, err
}

func GetBooks(client *ent.Client) ([]*ent.Book, error) {
	ctx := context.Background()
	books, err := client.Book.
		Query().
		All(ctx)
	return books, err
}
