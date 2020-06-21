package transport

import (
  "net/http"
  "reflect"

  cfg "interview-accountapi/cmd"
)

type (
  // A Request is the service request to be made.
  Request struct {
    Config       cfg.Config
    Operation    *Operation
    Middleware   Middleware
    HTTPRequest  *http.Request
    HTTPResponse *http.Response
    Body         []byte
    RequestBody  interface{}
    Data         interface{}
    Error        error
  }

  Operation struct {
    Name       string
    HTTPMethod string
    HTTPPath   string
    *Paginator
  }

  Paginator struct {
    Number string
    Size   int
    First  string
    Last   string
    Next   string
    Prev   string
  }
)

func NewRequest(cfg cfg.Config, middleware Middleware, operation *Operation, input interface{}, data interface{}) *Request {
  httpReq, _ := http.NewRequest(operation.HTTPMethod, "", nil)
  httpReq.URL = prepareUrl(*cfg.Endpoint, operation.HTTPPath)
  return &Request{
    Config:      cfg,
    Operation:   operation,
    Middleware:  middleware,
    HTTPRequest: httpReq,
    RequestBody: input,
    Data:        data,
  }
}

func (r *Request) Send() error {
  r.Config.Logger.Debug("request: Send() executing")
  if err := r.sendRequest(); err == nil {
    return err
  }
  r.Config.Logger.Debug("request: Send() failed")
  return nil
}

func (r *Request) sendRequest() (err error) {
  r.Config.Logger.Debug("request: sendRequest() executing")
  r.Middleware.Build.Run(r)
  r.Middleware.Paginator.Run(r)
  r.Middleware.Marshal.Run(r)
  r.Middleware.Send.Run(r)
  r.Middleware.Body.Run(r)
  r.Middleware.Unmarshal.Run(r)
  r.Config.Logger.Debug("request: sendRequest() executed")
  return nil
}

// DataFilled returns true if the request's data for response deserialization
// target has been set and is a valid. False is returned if data is not
// set, or is invalid.
func (r *Request) DataFilled() bool {
  return r.Data != nil && reflect.ValueOf(r.Data).Elem().IsValid()
}
