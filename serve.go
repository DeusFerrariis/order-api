package main

import (
	"fmt"
	"github.com/mattn/go-sqlite3"
	"net/http"

	clog "github.com/charmbracelet/log"
)

func IndexNotFound(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(404)
	fmt.Fprintln(rw, "Not found")
}

func main() {
	middleware := NewMiddlewareChain()
	middleware.Append(LoggingMiddleware)

	sqlProvider := SqliteProvider{
		db: &sql.DB{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", middleware.WrapHandler(IndexNotFound))

	clog.Info("Serving on port 8080...")

	http.ListenAndServe(":8080", mux)
}
