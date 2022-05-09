package yee

import (
	"fmt"
	"net/http"
)

// 请求处理函数
// todo:为什么request用指针而response用值，有啥讲究？
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET请求
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

// POST请求
func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	//engine需要实现了SeverHTTP方法
	return http.ListenAndServe(addr, engine)
}

// 只要实现了ServerHTTP方法就将对象传递给http.ListenAndServe()函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
