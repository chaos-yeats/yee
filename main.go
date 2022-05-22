package main

import (
	"log"
	"net/http"
	"time"

	"yee/yee"
)

func onlyForV2() yee.HandlerFunc {
	return func(c *yee.Context) {
		t := time.Now()
		c.Fail(500, "Internel Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := yee.New()
	r.Use(yee.Logger())

	r.Get("/", func(c *yee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v2 := r.Group("/v2")
	// curl -XPOST "http://127.0.0.1:9999/v2/login" -d "username=hello&password=world"
	v2.Use(onlyForV2())
	{
		v2.Get("/hello/:name", func(c *yee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
