package cmd

import (
  "net/http"
  "time"
)

const (
  Version  = "v1"
  Scheme   = "http"
  Domain   = "localhost"
  Port     = "8080"
  StatusOk = http.StatusOK
  // Timeout  = time.Second * 5
)

// https://github.com/aws/aws-sdk-go/blob/0ec921513a4536a315d4dd89416fc746b786ca2e/private/model/api/api.go

type Config struct {
  // An optional endpoint URL (hostname only or fully qualified URI)
  Endpoint *string

  // The logger writer interface to write logging messages to. Defaults to
  // standard out.
  Logger Logger

  // Timeout specifies a time limit for requests made by this
  // Client.
  //
  // A Timeout of zero means no timeout.
  //
  Timeout *time.Duration

  // The HTTP client to use when sending requests. Should defaults to `transport.DefaultClient`.
  HTTPClient *http.Client
}

func NewConfig() *Config {
  return &Config{}
}

// WithEndpoint sets a config Endpoint value returning a Config pointer
func (c *Config) WithEndpoint(endpoint string) *Config {
  c.Endpoint = &endpoint
  return c
}

// WithHTTPClient sets a config HTTPClient value returning a Config pointer
func (c *Config) WithHTTPClient(client *http.Client) *Config {
  c.HTTPClient = client
  return c
}

// WithLogger sets a config Logger value returning a Config pointer for
// chaining.
func (c *Config) WithLogger(logger Logger) *Config {
  c.Logger = logger
  return c
}
