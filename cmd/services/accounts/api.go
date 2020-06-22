package accounts

import (
  "fmt"
  net "net/http"

  "interview-accountapi/cmd/transport"
)

const name = "accounts"

type Api interface {
  fetch(id string) (out *Account, err error)
  list() (out *ResponsePaginate, err error)
  paginate(number string, size int) (out *ResponsePaginate, err error)
  delete(id string, version int) error
  create(a *Account) (*Account, error)
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

func (c *Service) fetch(id string) (out *Account, err error) {
  c.GetLogger().Debug("accounts-api: fetch() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   fmt.Sprintf("/organisation/accounts/%s", id),
  }

  data := &Response{}
  req := c.NewRequest(op, nil, data)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  out = &data.Data
  c.GetLogger().Debug("accounts-api: fetch() executing")
  return
}

func (c *Service) list() (out *ResponsePaginate, err error) {
  c.GetLogger().Debug("accounts-api: list() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   "/organisation/accounts",
  }

  out = &ResponsePaginate{}
  req := c.NewRequest(op, nil, &out)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  c.GetLogger().Debug("accounts-api: list() executing")
  return out, nil
}

func (c *Service) paginate(number string, size int) (out *ResponsePaginate, err error) {
  c.GetLogger().Debug("accounts-api: list() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodGet,
    HTTPPath:   "/organisation/accounts",
    Paginator:  &transport.Paginator{Number: number, Size: size},
  }

  out = &ResponsePaginate{}
  req := c.NewRequest(op, nil, out)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  c.GetLogger().Debug("accounts-api: list() executing")
  return out, nil
}

func (c *Service) create(a *Account) (*Account, error) {
  c.GetLogger().Debug("accounts-api: create() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodPost,
    HTTPPath:   "/organisation/accounts",
  }

  out := &Response{}
  req := c.NewRequest(op, &a, out)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  c.GetLogger().Debug("accounts-api: create() executing")
  return &out.Data, nil
}

func (c *Service) delete(id string, version int) error {
  c.GetLogger().Debug("accounts-api: delete() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: net.MethodDelete,
    HTTPPath:   fmt.Sprintf("/organisation/accounts/%s?version=%d", id, version),
  }

  req := c.NewRequest(op, nil, nil)
  if err := req.Send(); err != nil {
    req.Error = err
    return err
  }
  c.GetLogger().Debug("accounts-api: delete() executed")
  return nil
}
