package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	authur string
	price  float64
}

func main() {
	db, err := sql.Open("postgres", "postgres://henfo:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to Db", err.Error())
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging Db", err.Error())
		return
	}

	fmt.Println("Connected to Db!")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Error fecthing books", err.Error())
		return
	}

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.authur, &bk.price)
		if err != nil {
			fmt.Println("Error scanning book", err.Error())
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error scanning books", err.Error())
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, %.2f\n", bk.isbn, bk.title, bk.authur, bk.price)
	}

}
