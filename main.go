package main

import (
	"fmt"
	"net/http"

	"yee/yee"
)

func main() {
	fmt.Println("start http server...")
	r := yee.New()

	r.Get("/", func(c *yee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *yee.Context) {
		c.JSON(http.StatusOK, yee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
