package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
}

func main() {
	http.HandleFunc("/", handler)
	go func() {
		for {
			log.Println("Hello from within an infinite loop!")
			// Don't reach the disk quota too soon.
			time.Sleep(time.Millisecond)
		}
	}()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
