package db

import (
	"log"

	"api/ent"

	_ "github.com/go-sql-driver/mysql"
)

func Main() (*ent.Client) {
    client, err := ent.Open("mysql", "%:password@tcp(localhost:3306)/db/?parseTime=True")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    return client
}