package health

import (
  "net/http"

  "interview-accountapi/cmd/transport"
)

const name = "health-check"

type (
  Api interface {
    isUp() (out *CheckHealthOutput, err error)
  }

  CheckHealthOutput struct {
    Status string `json:"status"`
  }
)

// TODO: documentation
func (c *Service) isUp() (out *CheckHealthOutput, err error) {
  c.GetLogger().Debug("health-api: isUP() executing")
  op := &transport.Operation{
    Name:       name,
    HTTPMethod: http.MethodGet,
    HTTPPath:   "/health",
  }

  out = &CheckHealthOutput{}
  req := c.NewRequest(op, nil, out)
  if err := req.Send(); err != nil {
    return nil, err
  }
  c.GetLogger().Debug("health-api: IsUP() executed")
  return
}
