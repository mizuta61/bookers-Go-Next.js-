package db

import (
	"context"
	"log"

	"api/ent"

	_ "github.com/go-sql-driver/mysql"
)

func Main() (*ent.Client) {
    client, err := ent.Open("mysql", "root:password@tcp(db:3306)/mysql?parseTime=True")
    if err != nil {
        log.Fatal(err)
    }
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    return client
}