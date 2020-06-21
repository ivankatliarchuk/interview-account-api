package transport

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_ShouldPrepareUrl(t *testing.T) {
  fixtures := []struct {
    base     string
    path     string
    expected string
    error    bool
  }{
    {
      base:     "localhost",
      path:     "/v1/api",
      expected: "localhost/v1/api",
      error:    false,
    },
    {
      base:     "localhost:8090",
      path:     "/v1/api",
      expected: "localhost:8090/v1/api",
      error:    false,
    },
  }
  for _, fixture := range fixtures {
    url := prepareUrl(fixture.base, fixture.path)
    assert.Equal(t, url.String(), fixture.expected)
  }
}
