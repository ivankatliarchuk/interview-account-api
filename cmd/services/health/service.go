package health

import (
  "interview-accountapi/cmd/transport"
)

const (
  ServiceName = "health"        // Name of service.
)

type Service struct {
  transport.Session
}

// Creates a new instance of the HealthService check client.
func New(t transport.Session) *Service {
  return &Service{
    t,
  }
}

func (c *Service) IsUP() bool {
  c.GetLogger().Debug("health-api: IsUP() executing")
  _, err := c.isUp()
  return err == nil
}
