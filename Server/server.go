package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	ans, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println(ans, err)

}

func main() {
	http.HandleFunc("/", handler)
	timeAfter()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func timeAfter() {
	log.Println("some text is here")
	time.AfterFunc(3*time.Second, timeAfter)

}
