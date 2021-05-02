package main

import (
	"database/sql"

	db2 "github.com/felipebergamin/go-hexagonal/adapters/db"
	"github.com/felipebergamin/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Example Product", 30)
	productService.Enable(product)
}
