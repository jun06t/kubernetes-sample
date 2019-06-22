package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	cpus := runtime.NumCPU()
	fmt.Println("CPUs:", cpus)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.ListenAndServe(":8000", nil)
}
