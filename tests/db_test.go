package tests

import (
	"testing"

	"product-claims/db"

	"github.com/joho/godotenv"
)

func setup() {
	godotenv.Load("../.env")
	db.Connect()
}

func TestDBConnection(t *testing.T) {
	setup()

	err := db.Conn.Ping()
	if err != nil {
		t.Fatalf("db ping failed: %v", err)
	}

	t.Log("db connection works!")
}

func TestMaggiProductExists(t *testing.T) {
	setup()

	var title string
	err := db.Conn.QueryRow(
		"SELECT title FROM products WHERE id = $1",
		"cfe6aa75-5da8-44f5-b587-56857841ad9f",
	).Scan(&title)

	if err != nil {
		t.Fatalf("maggi product not found: %v", err)
	}

	t.Logf("found product: %s", title)
}