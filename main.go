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

		w.Write([]byte(fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr)))
		
		w.Write([]byte("RemoteAddr:\n"))
		for _, h := range r.Header {
			w.Write([]byte(fmt.Sprintf("%v", h)))
		}
		w.Write([]byte{'\n'})
		
		r.ParseForm()
		w.Write([]byte("Form Data:\n"))
		for _, f := range r.Form {
			w.Write([]byte(fmt.Sprintf("%v", f)))
		}
	})

	s := http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", port), Handler: r}
	s.ListenAndServe()
}
