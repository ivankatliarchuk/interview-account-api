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
  out, err := c.fetch(id)
  if err != nil {
    return nil, err
  }
  return out, nil
}

func (c *Service) Create(a *Account) (*Account, error) {
  c.GetLogger().Debug("accounts-service: Create() executing")
  result, err := c.create(a)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (c *Service) Delete(id string, version int) error {
  c.GetLogger().Debug("accounts-service: Delete() executing")
  err := c.delete(id, version)
  return err
}

func (c *Service) List() ([]Account, error) {
  c.GetLogger().Debug("accounts-service: Paginate() executing")
  result, err := c.list()
  if err != nil {
    return nil, err
  }
  return result.Data, nil
}

func (c *Service) Paginate(number string, size int) ([]Account, error) {
  c.GetLogger().Debug("accounts-service: Paginate() executing")
  result, err := c.paginate(number, size)
  if err != nil {
    return nil, err
  }
  return result.Data, nil
}
