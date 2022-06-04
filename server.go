package main

import (
	"log"
	"net/http"
	"os"

	"example.com/billy/graph"
	"example.com/billy/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

type Env struct {
	db *sqlx.DB
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sqlx.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln(err)
	}

	initDB(db)

	resolver := &graph.Resolver{DB: db}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB(db *sqlx.DB) {
	db.MustExec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, text varchar(255), completed boolean, userid TEXT)")
	db.MustExec("CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY, name varchar(255))")
}
