package transport

import (
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"

  "github.com/stretchr/testify/assert"

  cfg "interview-accountapi/cmd"
)

func Test_ShouldSetUrlWithPaginateOption(t *testing.T)  {
  fixtures := []struct {
    method string
    number string
    size int
    expected string
  }{
    {
      method: http.MethodGet,
      number: "4",
      size: 3,
      expected: "page[number]=4&page[size]=3",
    },
    {
      method: http.MethodGet,
      number: "last",
      size: 100,
      expected: "page[number]=last&page[size]=100",
    },
    {
      method: http.MethodPost,
      number: "last",
      size: 100,
      expected: "",
    },
  }
  for _, fixture := range fixtures {
    req := httptest.NewRequest(fixture.method, "http://example.com/foo", nil)
    r := &Request{
      HTTPRequest: req,
      Operation:&Operation{ Paginator: &Paginator{Number: fixture.number, Size: fixture.size}},
    }
    Paginate(r)
    actual := r.HTTPRequest.URL.Query()
    assert.Equal(t, actual.Encode(), url.PathEscape(fixture.expected))
  }
}

func Test_ShouldSetRightHeaders(t *testing.T)  {
  fixtures := []struct {
    method string
    expected http.Header
  }{
    {
      method:   http.MethodGet,
      expected: http.Header{
        "Accept":[]string{"application/vnd.api+json"},
        "User-Agent": []string{"accounts-api-sdk-go"}},
    },
    {
      method:   http.MethodDelete,
      expected: http.Header{
        "Accept":[]string{"application/vnd.api+json"},
        "User-Agent": []string{"accounts-api-sdk-go"}},
    },
    {
      method:   http.MethodPost,
      expected: http.Header{
        "Accept":[]string{"application/vnd.api+json"},
        "Content-Type":[]string{"application/vnd.api+json"},
        "User-Agent": []string{"accounts-api-sdk-go"}},
    },
  }
  for _, fixture := range fixtures {
    req := httptest.NewRequest(fixture.method, "http://localhost", nil)
    r := &Request{
      Config: cfg.Config{Logger: &MockLogger{}},
      Operation:  &Operation{HTTPMethod: fixture.method},
      HTTPRequest: req,
    }
    SetBaseHeadersMiddleware(r)
    actual := r.HTTPRequest.Header
    assert.Equal(t, actual, fixture.expected)
  }
}
