package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/victornm/todo/models/todo"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := todo.New(r.URL.Path[1:], r.URL.Path[1:])
	log.Println("Received request")
	fmt.Fprintf(w, "<h1>%s</h1><div>%d %s</div>", t.Title, t.ID, t.Detail)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
