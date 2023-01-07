package core

import (
	"github.com/B1NARY-GR0UP/inquisitor/core"
	"net/http"
	"strings"
)

// RouterGroup must implement IRouter
var _ IRouter = (*RouterGroup)(nil)

type IRouter interface {
	IRoute
	GROUP(string, ...HandlerFunc) *RouterGroup
}

type IRoute interface {
	USE(...HandlerFunc)

	// GET POST PUT DELETE HTTP request methods
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
}

type RouterGroup struct {
	engine   *Engine
	Handlers HandlersChain
	basePath string
	isRoot   bool
}

func (rg *RouterGroup) GROUP(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: handlers,
		basePath: rg.calculateAbsolutePath(relativePath),
		engine:   rg.engine,
	}
}

func (rg *RouterGroup) USE(middleware ...HandlerFunc) {
	rg.Handlers = append(rg.Handlers, middleware...)
}

func (rg *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodGet, relativePath, handlers)
}

func (rg *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodPost, relativePath, handlers)
}

func (rg *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodPut, relativePath, handlers)
}

func (rg *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodDelete, relativePath, handlers)
}

func (rg *RouterGroup) handle(method, relativePath string, handlers HandlersChain) {
	absolutePath := rg.calculateAbsolutePath(relativePath)
	mergedHandlers := rg.combineHandlers(handlers)
	// note that the relative path is changed to absolutePath and handlers are changed to mergedHandlers
	rg.engine.addRoute(method, absolutePath, mergedHandlers)
}

func (rg *RouterGroup) calculateAbsolutePath(relativePath string) string {
	if relativePath == "" {
		return rg.basePath
	}
	sb := &strings.Builder{}
	if rg.basePath != "/" {
		sb.WriteString(rg.basePath)
	}
	sb.WriteString(relativePath)
	return sb.String()
}

func (rg *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	totalSize := len(rg.Handlers) + len(handlers)
	core.Infof("Number of handlers: %v", totalSize)
	mergedHandlers := make(HandlersChain, totalSize)
	copy(mergedHandlers, rg.Handlers)
	copy(mergedHandlers[len(rg.Handlers):], handlers)
	return mergedHandlers
}
