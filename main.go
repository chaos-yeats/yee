package main

import (
	"net/http"
	"yee/yee"
)

func main() {
	r := yee.Default()
	r.Get("/", func(c *yee.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})
	// index out of range for testing Recovery()
	r.Get("/panic", func(c *yee.Context) {
		names := []string{"yee"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
