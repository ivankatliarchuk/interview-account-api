package transport

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "reflect"
)

type RequestData struct {
  Data interface{} `json:"data"`
}

// Marshall/Unmarshall provides the interface for unmarshalling a payload's
// into a SDK shape.
var (
  MarshalMiddleware   = NamedMiddleware{Name: "middleware::marshal-json", Fn: Marshal}
  UnmarshalMiddleware = NamedMiddleware{Name: "middleware::unmarshal-json", Fn: Unmarshal}
)

func Marshal(r *Request) {
  if r.Operation.HTTPMethod == http.MethodPost && r.RequestBody != nil {
    r.Config.Logger.Debug("marshall: Marshal() ", reflect.TypeOf(r.RequestBody))
    if err := marshal(r, r.RequestBody); err != nil {
      r.Error = err
    }
  }
}

func marshal(r *Request, data interface{}) error {
  var buf bytes.Buffer
  err := json.NewEncoder(&buf).Encode(&RequestData{data})
  if err != nil {
    return err
  }
  r.HTTPRequest.Body = ioutil.NopCloser(&buf)
  return nil
}

// Unmarshal unmarshals a response body for the REST JSON protocol.
func Unmarshal(r *Request) {
  r.Config.Logger.Debug("unmarshall: Unmarshal() ", reflect.TypeOf(r.Data))
  if r.DataFilled() {
    if err := unmarshal(r, r.Data); err != nil {
      r.Error = err
    }
  }
}

func unmarshal(r *Request, data interface{}) error {
  if r.Body != nil && len(r.Body) > 0 {
    if err := json.Unmarshal(r.Body, data); err != nil {
      return err
    }
  }
  return nil
}
