package db

import (
	"context"
	"log"
	"os"

	"api/ent"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMariadb() (*ent.Client) {
    db_client := os.Getenv("DB_CLIENT")
    db_url := os.Getenv("DB_URL")
    client, err := ent.Open(db_client, db_url)
    if err != nil {
        log.Fatal(err)
    }
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    return client
}