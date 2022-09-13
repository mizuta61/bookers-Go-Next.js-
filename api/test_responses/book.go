package test_responses

import (
	"api/model"
	"time"
)	

func TestBooks() []model.Book {
	books := [] model.Book{
		{
      Id: 1,
      Title: "test1",
      Body: "test1",
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    },
    {
      Id: 2,
      Title: "test2",
      Body: "test2",
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    },
	}
	return books
}