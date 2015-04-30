package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.String())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

func main() {
	args := os.Args[1:]
	dir, err := os.Getwd()
	if len(args) > 0 {
		dir, err = filepath.Abs(args[0])
		if err != nil {
			dir, _ = os.Getwd()
		}
	}
	port := ":8080"
	if len(args) > 1 {
		port = fmt.Sprintf(":%v", args[1])
	}

	fmt.Printf("Serving directory %v over HTTP on port %v...\n", dir, port[1:])
	panic(http.ListenAndServe(port, loggingHandler(http.FileServer(http.Dir(dir)))))
}
