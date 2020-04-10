package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/leandroribeiro/go-labs/api-lab2/home"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "go-labs ", log.LstdFlags | log.Lshortfile)

	database, err := DatabaseInitialize()

	h := home.NewHandlers(logger, database)

	mux := http.NewServeMux()
	h.SetUpRoutes(mux)

	logger.Println("server starting")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func DatabaseInitialize() (*sqlx.DB, error) {

	database, err := sqlx.Open("sqlite3", "./demo.db")

	if err != nil {
		log.Fatalln(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", "Raboy")
	return database, err
}

