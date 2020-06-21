package transport

import (
  "interview-accountapi/cmd"
)

type (
  // A Client implements the base client request and response handling
  // used by all service clients.
  Client struct {
    Config     cmd.Config
    Middleware Middleware
  }

  Session interface {
    GetConfig() cmd.Config
    GetLogger() cmd.Logger
    GetMiddleware() Middleware
    NewRequest(operation *Operation, input interface{}, data interface{}) *Request
  }
)

// Will return a pointer to a new initialized default service client.
func DefaultClient() *Client {
  p := Preconfigure()
  svc := &Client{
    Config:     p.Config,
    Middleware: p.Middleware,
  }
  return svc
}

func (c *Client) GetConfig() cmd.Config {
  return c.Config
}

func (c *Client) GetMiddleware() Middleware {
  return c.Middleware
}

func (c *Client) GetLogger() cmd.Logger {
  return c.GetConfig().Logger
}

// NewRequest returns a new Request pointer with operation and parameters.
func (c *Client) NewRequest(operation *Operation, input interface{}, data interface{}) *Request {
  c.GetLogger().Debug("client: NewRequest() executing")
  return NewRequest(c.GetConfig(), c.GetMiddleware(), operation, input, data)
}
