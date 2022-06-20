package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// executes the query - based on the context of the request
func slowQuery(cxt context.Context) error {

	// execute the query of simulating a 5 second sleep of a slow query
	_, err := db.ExecContext(cxt, "SELECT pg_Sleep(5)")
	return err
}

// Handler for the request http://localhost:8000
func slowHandler(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	// pass the context for slowQuery to execute
	if err := slowQuery(r.Context()); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Printf("Warning %s\n", err.Error())
		default:
			log.Printf("Error %s\n", err.Error())
		}
		return
	}

	fmt.Println(w, "ok")
	fmt.Printf("SlowHandler took %s to complete. \n", time.Since(start))
}

func main() {

	var err error

	// 1. create a connectionstring to the database wonderland via the user alice
	connstr := "host=localhost port=5432 user=alice password=pa$$word dbname=wonderland sslmode=disable"

	// 2. open the database
	if db, err = sql.Open("postgres", connstr); err != nil {
		log.Fatal(err)
	}

	// 3. create the context for the request that will timeout within 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 4. use the context to allow the database to be pinged
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	// 5. create the server configuration to handle requests
	srv := http.Server{
		Addr:         "localhost:8000",                                                                // provide the address of the server
		WriteTimeout: 2 * time.Second,                                                                 //set the duration of the timeout
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!!!"), // set the handler to the request and the duration to return a timeout
	}

	if err = srv.ListenAndServe(); err != nil {
		fmt.Printf("Server ended due to: %s. \n", err.Error())
	}
}

// [devworks@] $: psql -U postgres
// postgres=# \l
// postgres=# create database alice;
// postgres=# \l
// postgres=# alter database alice rename to wonderland;
// postgres=# create user alice with encrypted password 'pa$$word';
// postgres=# grant all privileges on database wonderland to alice;
