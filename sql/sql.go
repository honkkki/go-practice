package main

import (
	"database/sql"
	"log"
)

func main()  {
	_, err := sql.Open("mysql", "")

	if err != nil {
		log.Fatal("error")
	}

}
