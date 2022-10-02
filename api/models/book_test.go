package models

import (
	"api/ent"
	"api/ent/enttest"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestBook(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	book := &ent.Book{
		Title: "test1",
		Body:  "test1",
	}
	created, err := CreateBook(client, book)
	if err == nil {
		log.Println("book was created: ", created)
	} else {
		log.Fatal(err)
	}
	t.Log("TestCreateBooks成功")
	book_id := 1
	book, err = GetBook(client, book_id)
	if err == nil {
		log.Println("book was gotten: ", book)
	} else {
		log.Fatal(err)
	}
	t.Log("TestGetBook成功")
	books, err := GetBooks(client)
	if err == nil {
		log.Println("all_books: ", books)
	} else {
		log.Fatal(err)
	}
	t.Log("TestGetBooks成功")
	form := ent.Book{
		Title: "test1update",
		Body:  "test1update",
	}
	book, err = UpdateBook(client, book_id, form)
	if err == nil {
		log.Println("book was updated: ", book)
	} else {
		log.Fatal(err)
	}
	t.Log("TestUpdateBook成功")
	book, err = GetBook(client, book_id)
	if err == nil {
		log.Println("book was gotten: ", book)
	} else {
		log.Fatal(err)
	}
	t.Log("Update確認")
  err = DestroyBook(client, book_id)
	if err != nil {
		log.Fatal(err)
	} 
  books, err = GetBooks(client)
	if err == nil {
		log.Println("all_books: ", books)
	} else {
		log.Fatal(err)
	}
	t.Log("TestDelete成功")
}

