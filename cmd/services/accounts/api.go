package accounts

import (
  "fmt"
  net "net/http"

  "interview-accountapi/cmd/transport"
)

const name = "accounts"

type Api interface {
  fetch(id string) (req *transport.Request, out *Response)
  list() (req *transport.Request, out *ResponsePaginate)
  paginate(number string, size int) (req *transport.Request, out *ResponsePaginate)
  delete(id string, version int) (req *transport.Request)
  create(a *Account) (req *transport.Request, out *Response)
}

type (
  Response struct {
    Data  Account `json:"data"`
    Links Links   `json:"links"`
  }

  ResponsePaginate struct {
    Data  []Account `json:"data"`
    Links Links     `json:"links"`
  }

  Links struct {
    First string `json:"first"`
    Last  string `json:"last"`
    Self  string `json:"self"`
    Next  string `json:"next"`
    Prev  string `json:"prev"`
  }
)

func (c *Service) fetch(id string) (req *transport.Request, out *Response) {
  c.GetLogger().Debug("accounts-api: fetch() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   fmt.Sprintf("/organisation/accounts/%s", id),
  }

  out = &Response{}
  req = c.NewRequest(op, nil, out)
  c.GetLogger().Debug("accounts-api: fetch() executing")
  return
}

func (c *Service) list() (req *transport.Request, out *ResponsePaginate) {
  c.GetLogger().Debug("accounts-api: list() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   "/organisation/accounts",
  }

  out = &ResponsePaginate{}
  req = c.NewRequest(op, nil, out)
  c.GetLogger().Debug("accounts-api: list() executing")
  return
}

func (c *Service) paginate(number string, size int) (req *transport.Request, out *ResponsePaginate) {
  c.GetLogger().Debug("accounts-api: list() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   "/organisation/accounts",
    Paginator:  &transport.Paginator{Number: number, Size: size},
  }

  out = &ResponsePaginate{}
  req = c.NewRequest(op, nil, out)
  c.GetLogger().Debug("accounts-api: list() executing")
  return
}

func (c *Service) create(a *Account) (req *transport.Request, out *Response) {
  c.GetLogger().Debug("accounts-api: create() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodPost,
    HTTPPath:   "/organisation/accounts",
  }

  out = &Response{}
  req = c.NewRequest(op, &a, out)
  c.GetLogger().Debug("accounts-api: create() executing")
  return
}

func (c *Service) delete(id string, version int) (req *transport.Request) {
  c.GetLogger().Debug("accounts-api: delete() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodDelete,
    HTTPPath:   fmt.Sprintf("/organisation/accounts/%s?version=%d", id, version),
  }

  req = c.NewRequest(op, nil, nil)
  c.GetLogger().Debug("accounts-api: delete() executing")
  return
}
