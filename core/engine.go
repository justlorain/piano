package core

import (
	"github.com/B1NARY-GR0UP/inquisitor/core"
	"net/http"
	"strings"
	"sync"
)

type Engine struct {
	// RouterGroup is a composition of Engine so that Engine can use RouterGroup functions
	RouterGroup

	forest    MethodForest
	options   *Options
	ctxPool   sync.Pool
	maxParams uint16
}

func NewEngine(opts *Options) *Engine {
	e := &Engine{
		forest: make(MethodForest, 0, 5),
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			isRoot:   true,
		},
		options: opts,
	}
	e.RouterGroup.engine = e
	// TODO: Why, how maxParams assigned??
	e.ctxPool.New = func() any {
		return e.allocateContext(e.maxParams)
	}
	return e
}

// Play Start the Server
func (e *Engine) Play() {
	core.Info("Server Start")
	err := http.ListenAndServe(e.options.Addr, e)
	if err != nil {
		panic("Server Start Failed")
	}
}

// ServeHTTP core function, replace DefaultServeMux
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pk := e.ctxPool.Get().(*PianoKey)
	// inject content to *PianoKey
	pk.Writer = w
	pk.Request = req
	pk.refresh()
	e.handleHTTPRequest(pk)
	e.ctxPool.Put(pk)
}

func (e *Engine) handleHTTPRequest(pk *PianoKey) {
	for _, tree := range e.forest {
		if tree.method == pk.Request.Method {
			// TODO: match tree
			tree.matchBranch()
			// TODO: serve
		}
	}
	// TODO: use find and next to handler request
}

func (e *Engine) addRoute(method, path string, handlers HandlersChain) {
	isValid := validateRoute(method, path, handlers)
	if !isValid {
		panic("please check your route")
	}
	core.Infof("Register route: [%v] %v", strings.ToUpper(method), path)
	methodTree, ok := e.forest.get(method)
	// create a new method tree if no match in the forest
	if !ok {
		methodTree = &tree{
			method: method,
			root:   &node{},
		}
		e.forest = append(e.forest, methodTree)
	}
	methodTree.addRoute(path, handlers)
	// update mexParams
	if paramCount := calculateParam(path); paramCount > e.maxParams {
		e.maxParams = paramCount
	}
}

func (e *Engine) allocateContext(maxParams uint16) *PianoKey {
	return NewContext(maxParams)
}
