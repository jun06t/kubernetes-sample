package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	cpus := runtime.NumCPU()
	fmt.Println("CPUs:", cpus)

	for i := 0; i < 100; i++ {
		goroutine()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.ListenAndServe(":8000", nil)
}

func goroutine() {
	go func() {
		var counter int64
		for {
			counter++
		}
	}()
}
