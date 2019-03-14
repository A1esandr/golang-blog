package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var PORT = flag.Int("port", 15001, "listen port")

func Hello(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'name' is: " + string(key))

	message := "Hello, " + string(key) + "!"
	head := "<head><title>" + message + "</title></head>"
	response := "<html>" + head + "<body>" + message + "</body></html>"

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *PORT)
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
