package transport

// A Middleware provides a collection of middleware for various stages of requests flow.
type Middleware struct {
  Build     MiddlewareList
  Body      MiddlewareList
  Send      MiddlewareList
  Marshal   MiddlewareList
  Unmarshal MiddlewareList
  Paginator MiddlewareList
}

// A MiddlewareList manages zero or more handlers in a list.
type MiddlewareList struct {
  list []NamedMiddleware
  // Called after each middleware in the list is called
  AfterEachFn func(item MiddlewareRunItem) bool
}

// A NamedMiddleware is a struct that contains a name and function callback.
type NamedMiddleware struct {
  Name string
  Fn   func(*Request)
}

// A MiddlewareRunItem represents an entry in the MiddlewareList which is being run.
type MiddlewareRunItem struct {
  Index   int
  NamedMiddleware
  *Request
}

// Append named handler to the back of the middleware collection.
func (l *MiddlewareList) Append(n NamedMiddleware) {
  if cap(l.list) == 0 {
    l.list = make([]NamedMiddleware, 0, 5)
  }
  l.list = append(l.list, n)
}

// Run executes a handler in the list with a given request object.
func (l *MiddlewareList) Run(r *Request) {
  for i, h := range l.list {
    h.Fn(r)
    item := MiddlewareRunItem{
      Index: i, NamedMiddleware: h, Request: r,
    }
    item.Request.Config.Logger.Debug("middleware: run() ", h.Name)
    if l.AfterEachFn != nil && !l.AfterEachFn(item) {
      return
    }
  }
}
