package main

import (
	"log"
	"net/http"

	_ "crud/controllers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
