package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/kill", shutdown)
	http.ListenAndServe(":8081", nil)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func homePage(res http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(res, r)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Fprint(res, "Welcome")

}
