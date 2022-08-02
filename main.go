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
	r.HandleFunc("/adr", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Write([]byte(fmt.Sprintf("%s\n",r.RemoteAddr)))
		w.Write([]byte(fmt.Sprintf("%#v\n",r.Header)))
		w.Write([]byte(fmt.Sprintf("%#v\n",r.Form)))
	})

	s := http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", port), Handler: r}
	s.ListenAndServe()
}
