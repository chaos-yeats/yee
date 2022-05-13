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

	// curl 模拟表单：https://blog.csdn.net/freedomwjx/article/details/43278157
	// curl -XPOST "http://127.0.0.1:9999/login" -d "username=hello&password=world"
	r.Post("/login", func(c *yee.Context) {
		c.JSON(http.StatusOK, yee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
