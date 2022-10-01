package models

import (
	"api/ent"
	"api/ent/enttest"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestBook(t *testing.T) {
	// t.Setenv("DB_CLIENT", "sqlite3")
	// t.Setenv("DB_URL", "file:ent?mode=memory&cache=shared&_fk=1")
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	// ctx := context.Background()
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
	t.Log("TestGetBooks成功")
}

// func TestGetBook(t *testing.T) {
// 	t.Setenv("DB_CLIENT", "sqlite3")
// 	t.Setenv("DB_URL", "file:ent?mode=memory&cache=shared&_fk=1")
// 	ctx := context.Background()
// 	book_id := 1
// 	book, err := GetBook(ctx, book_id)
// 	if err == nil {
// 		log.Println("book was gotten: ", book)
// 	} else {
// 		log.Fatal(err)
// 	}
// 	t.Log("TestGetBooks終了")
// }
