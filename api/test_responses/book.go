package test_responses

import (
	"api/ent"
)

func TestBooks() []ent.Book {
	books := []ent.Book{
		{
			ID:    1,
			Title: "test1",
			Body:  "test1",
		},
		{
			ID:    2,
			Title: "test2",
			Body:  "test2",
		},
	}
	return books
}
