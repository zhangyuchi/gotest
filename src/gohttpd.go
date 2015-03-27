package main

import (
	"fmt"
	"http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
        runtime.GOMAXPROCS(4); // 4 cores
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
