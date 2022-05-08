package main

import (
	"fmt"
	"net/http"

	"yee/yee"
)

func main() {
	fmt.Println("start http server...")
	r := yee.New()

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Post("/world", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Method=%q\n", req.Method)
	})

	r.Run(":9999")
}
