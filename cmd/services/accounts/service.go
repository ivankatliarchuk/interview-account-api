package accounts

import (
  "interview-accountapi/cmd/transport"
)

type (
  AccountService interface {
    Fetch(id string) (*Account, error)
    Create(a *Account) (*Account, error)
    Delete(id string, version int) (err error)
    List() ([]Account, error)
    Paginate(number string, size int) ([]Account, error)
  }
  Service struct {
    transport.Session
  }
)

func New(t transport.Session) *Service {
  return &Service{
    t,
  }
}

func (c *Service) Fetch(id string) (*Account, error) {
  c.GetLogger().Debug("accounts-service: Fetch() executing")
  req, out := c.fetch(id)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  return &out.Data, nil
}

func (c *Service) Create(a *Account) (*Account, error) {
  c.GetLogger().Debug("accounts-service: Create() executing")
  req, out := c.create(a)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  return &out.Data, nil
}

func (c *Service) Delete(id string, version int) (err error) {
  c.GetLogger().Debug("accounts-service: Delete() executing")
  req := c.delete(id, version)
  if err := req.Send(); err != nil {
    req.Error = err
    return err
  }
  return nil
}

func (c *Service) List() ([]Account, error) {
  c.GetLogger().Debug("accounts-service: Paginate() executing")
  req, out := c.list()
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  return out.Data, nil
}

func (c *Service) Paginate(number string, size int) ([]Account, error) {
  c.GetLogger().Debug("accounts-service: Paginate() executing")
  req, out := c.paginate(number, size)
  if err := req.Send(); err != nil {
    req.Error = err
    return nil, err
  }
  return out.Data, nil
}
