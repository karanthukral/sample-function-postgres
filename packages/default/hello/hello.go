package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	fmt.Println("Connecting to postgres...")
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Pinging...")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	msg["body"] = fmt.Sprintf("%s. DB ping successful.", msg["body"])

	return msg
}
