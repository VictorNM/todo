package main

import (
	"github.com/victornm/todo/api"
	"log"
)

func main() {
	r := api.InitRouter()
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
