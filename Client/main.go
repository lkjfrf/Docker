package main

import (
	"log"
	"net/http"
)

func main() {
	request1()
}

func hoemPage(response http.ResponseWriter, r *http.Request) {
	log.Println("GO web api")
}

func aboueMe(response http.ResponseWriter, r *http.Request) {
	who := "songsong"

	log.Println(response, "AboueMe", who)
}

func request1() {
	http.HandleFunc("/", hoemPage)
	http.HandleFunc("/aboueme", aboueMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
