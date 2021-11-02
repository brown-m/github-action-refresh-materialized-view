package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatal("Unable to create database connection: ", err)
	}
	defer db.Close()

	fmt.Printf("Refreshing materialized view %s...\n", os.Getenv("INPUT_VIEW"))
	query := "refresh materialized view " + os.Getenv("INPUT_VIEW")

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Println("...DONE")
}
