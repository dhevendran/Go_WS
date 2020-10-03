package main

import (
	"database/sql"
	"fmt"
	"log"

	// we are importing this package so that it can register its drivers with the database/sql package,
	// and we use the _ identifier to tell Go that we still want this included even though we will never directly reference the package in our code.
	_ "github.com/lib/pq"
)

// PostgreSQL database info
// You set your username and password when you run:
// psql -U <user> <password>
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "emp"
)

func main() {
	//Just to log the errors
	fmt.Println("Main is working")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening a connection to the database
	// The sql.Open() function takes two arguments - a driver name, and a string that tells that driver how to connect to our database - and then returns a pointer to a sql.DB and an error.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db.Ping() forces our code to actually open up a connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", db)
	fmt.Println("Successfully connected!")

	sqlStatement := `
	INSERT INTO emp_table (role_id,first_name, last_name)
	VALUES ($1,$2, $3)
	RETURNING first_name`
	first_name := ""
	err = db.QueryRow(sqlStatement, "3", "Dhevendran", "Kulandaivelu").Scan(&first_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", first_name)

}
