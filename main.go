package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"yee/yee"
)

type Student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// 模版的使用方法：https://www.cnblogs.com/wanghui-garcia/p/10385062.html
func main() {
	r := yee.New()
	r.Use(yee.Logger())

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &Student{Name: "Chaos", Age: 20}
	stu2 := &Student{Name: "Yeats", Age: 22}

	r.Get("/", func(c *yee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Get("/students", func(c *yee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", yee.H{
			"title":  "Data transport",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})

	r.Get("/date", func(c *yee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", yee.H{
			"title": "Date show",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
