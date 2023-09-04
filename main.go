package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-react-todo/server/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}


Esta aqui e minha edicao