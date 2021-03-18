package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Go Web API")
	fmt.Printf("Endpoint Hit: homePage")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "Ahmed Mohamed"

	fmt.Fprintf(w, "A little bit about Ahmed Mohamed...")
	fmt.Println("Endpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	request1()
}
