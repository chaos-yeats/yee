package yee

import (
	"html/template"
	"net/http"
	"strings"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// 请求处理函数
// todo:为什么request用指针而response用值，有啥讲究？
type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine} // 神奇的操作
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGoup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGoup)

	return newGoup
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 添加路由
func (group *RouterGroup) addRouter(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, handler)
}

// GET请求
func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

// POST请求
func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}

func (group *RouterGroup) Run(addr string) (err error) {
	//engine需要实现了SeverHTTP方法
	return http.ListenAndServe(addr, group.engine)
}

// 只要实现了ServerHTTP方法就将对象传递给http.ListenAndServe()函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handler(c)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
