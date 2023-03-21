package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username = "aaaaaaaa"
	password = "ssssssss"
	endpoint = "maindb.czpld8fke1ht.us-east-1.rds.amazonaws.com"
	port     = "3306"
	db_name  = "maindb"
)

func main() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, endpoint, port, db_name)
	// db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	fmt.Println("Success!")
}
