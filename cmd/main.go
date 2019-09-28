package main

import (
	"github.com/victornm/todo/router"
	"log"
)

func main() {
	r := router.Init()
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
