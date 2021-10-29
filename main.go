package main

import (
	"fmt"
	"log"
	"net/http"
)

const HOST = "0.0.0.0"
const PORT = 5461

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server working")
	})

	fmt.Printf("Server started on %s:%d\n", HOST, PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), nil))
}
