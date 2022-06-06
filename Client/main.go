package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func hoemPage(response http.ResponseWriter, r *http.Request) {
	fmt.Println("GO web api")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "songsong"

	fmt.Fprintf(response, "hihi")
	fmt.Println("EndPoint Hit : ", who)
}

func request1() {
	http.HandleFunc("/", hoemPage)
	http.HandleFunc("/aboutme", aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
