package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	r := mux.NewRouter()
	r.HandleFunc("adr", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(r.RemoteAddr)) })
	s := http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", port), Handler: r}
	s.ListenAndServe()
}
