package transport

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"

  cfg "interview-accountapi/cmd"
)

func Test_ShouldMarshallFromJsonObject(t *testing.T) {
  r := &Request{Config: cfg.Config{Logger: &MockLogger{}}, Operation: &Operation{HTTPMethod: http.MethodGet}}
  Marshal(r)
  assert.Nil(t, r.HTTPRequest)

  Marshal(r)
  req := httptest.NewRequest("POST", "http://example.com/foo", nil)
  r = &Request{
    Config:      cfg.Config{Logger: &MockLogger{}},
    Operation:   &Operation{HTTPMethod: http.MethodPost},
    RequestBody: &MockStruct{Value: "12"},
    HTTPRequest: req,
  }
  Marshal(r)
  assert.Equal(t, fmt.Sprintf("%s", r.HTTPRequest.Body), "{{\"data\":{\"value\":\"12\"}}\n}")
}

func Test_ShouldNotUnMarshallJsonObject(t *testing.T)  {
  json := `{"data":"Test Name","full_name":"test","owner":"octocat"}`
  r := &Request{Config: cfg.Config{Logger: &MockLogger{}}, Body:[]byte(json)}
  Unmarshal(r)
  assert.Nil(t, r.Data)
}

func Test_ShouldMarshallJsonObject(t *testing.T)  {
  json := `{"value":"sdk"}`
  out := &MockStruct{}
  r := &Request{Config: cfg.Config{Logger: &MockLogger{}}, Body:[]byte(json), Data: out}
  Unmarshal(r)
  assert.NotNil(t, r.Data)
  assert.Equal(t, out.Value, "sdk")
}
