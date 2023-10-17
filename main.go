package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NamozovAzizbek/go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8000....")

	log.Fatal(http.ListenAndServe(":8000", r))
}
