package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const RETRIES = 3

func main() {
	failures := 0

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatal("Unable to create database connection: ", err)
	}
	defer db.Close()

	fmt.Printf("Refreshing materialized view %s...\n", os.Getenv("INPUT_VIEW"))
	query := fmt.Sprintf("refresh materialized view %s;", os.Getenv("INPUT_VIEW"))

	for failures < RETRIES {
		_, err = db.Exec(query)
		if err != nil {
			fmt.Printf("Failed to execute query: %s\n", err)
			failures += 1
		} else {
			fmt.Println("...DONE")
			break
		}
	}

	if failures == RETRIES-1 {
		os.Exit(1) // Exit with non-zero value to indicate an error
	}
}
