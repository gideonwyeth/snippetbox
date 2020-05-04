package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gideonwyeth/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

// application struct for holding the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	// server address flag
	addr := flag.String("addr", ":8080", "HTTP network address")
	// db dsn flag
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create db connection pool
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// initialize a new instance of web application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	// create server with custom error logging
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start server
	infoLog.Printf("Starting server on %s\n", *addr)
	errorLog.Fatalln(srv.ListenAndServe())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
