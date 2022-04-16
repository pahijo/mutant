package main

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

	"flag"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load("enviroment.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func Create(w http.ResponseWriter, m bool) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var analisis Analisis
	analisis.Mutant = m

	insertID := insertAnalisis(analisis)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Dna created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// insert one user in the DB
func insertAnalisis(dna Analisis) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	sqlStatement := `INSERT INTO analisis (mutant) VALUES ($1) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, dna.Mutant).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getRatio() (XDna, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var item XDna

	// create the select sql query
	sqlStatement := `select count(*) as humano,(select count(*) as mutante from analisis where mutant = true) from analisis	where mutant = false`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {

		// unmarshal the row object to user
		err = rows.Scan(&item.Mutante, &item.Humano)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
	}
	mutant := flag.Int64("Mutante", item.Mutante, "total number of mutants")
	human := flag.Int64("Humano", item.Humano, "number of humans")

	item.RatioT = NewRatio(*mutant, *human)
	// return empty user on error
	return item, err
}

func (r Ratio) String() string {
	return fmt.Sprintf("%.0f%%", r)
}

func NewRatio(mutante, humano int64) Ratio {
	return Ratio(mutante) / Ratio(humano) * 100.0
}
