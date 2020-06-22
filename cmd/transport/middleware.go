package transport

import (
  "io"
  "io/ioutil"
  "net/http"
  "strconv"
)

var (
  UnmarshalDiscardBodyMiddleware = NamedMiddleware{Name: "middleware::unmarshal-discard-body", Fn: UnmarshalDiscardBody}
  PaginatorMiddleware            = NamedMiddleware{Name: "middleware::request-paginate", Fn: Paginate}
  DoRequestMiddleware            = NamedMiddleware{Name: "middleware::request-do", Fn: DoRequest}
  CopyResponseBodyMiddleware     = NamedMiddleware{Name: "middleware::response-body", Fn: CopyResponseBody}
  BaseHeadersMiddleware          = NamedMiddleware{Name: "middleware::request-base-headers", Fn: SetBaseHeadersMiddleware}
  // NoBody is a transport.NoBody reader instructing Go HTTP client to not include body in the HTTP request.
  NoBody = http.NoBody
)

// UnmarshalDiscardBody is a request handler to empty a response's body and closing it.
func UnmarshalDiscardBody(r *Request) {
  if r.HTTPResponse == nil || r.HTTPResponse.Body == nil {
    r.Config.Logger.Debug("middleware: UnmarshalDiscardBody() values empty")
    return
  }
  _, _ = io.Copy(ioutil.Discard, r.HTTPResponse.Body)
  _ = r.HTTPResponse.Body.Close()
}

func Paginate(r *Request) {
  if r.HTTPRequest.Method == http.MethodGet && r.Operation.Paginator != nil {
    url := r.HTTPRequest.URL.Query()
    url.Add("page[number]", r.Operation.Number)
    url.Add("page[size]", strconv.Itoa(r.Operation.Size))
    r.HTTPRequest.URL.RawQuery = url.Encode()
  }
}

func DoRequest(r *Request) {
  send := sendRequest
  // Strip off the request body if the NoBody reader was used as a
  // place holder for a request body.
  if NoBody == r.HTTPRequest.Body {
    // Use a shallow copy of the transport.Request to ensure the race condition
    // of transport on Body will not trigger
    reqOrig, reqCopy := r.HTTPRequest, *r.HTTPRequest
    reqCopy.Body = nil
    r.HTTPRequest = &reqCopy
    defer func() {
      r.HTTPRequest = reqOrig
    }()
  }
  var err error
  r.HTTPResponse, err = send(r)
  r.Config.Logger.Debug("middleware: DoRequest executed")
  if err != nil {
    // TODO: proper implement
    panic(err)
  }
}

func CopyResponseBody(r *Request) {
  if r.HTTPResponse == nil || r.HTTPResponse.Body == nil {
    r.Config.Logger.Debug("middleware: CopyResponseBodyMiddleware values empty")
    return
  }
  body, err := ioutil.ReadAll(r.HTTPResponse.Body)
  if err != nil {
    panic(err)
  }
  r.Body = body
}

func SetBaseHeadersMiddleware(r *Request) {
  r.Config.Logger.Debug("middleware: RequestSetHeadersMiddleware set headers.", r.Operation.HTTPMethod)
  r.HTTPRequest.Header.Set("User-Agent", "accounts-api-sdk-go")
  switch r.Operation.HTTPMethod {
  case
    http.MethodGet,
    http.MethodDelete:
    r.HTTPRequest.Header.Set("Accept", "application/vnd.api+json")
    return
  case
    http.MethodPost,
    http.MethodPut:
    r.HTTPRequest.Header.Set("Accept", "application/vnd.api+json")
    r.HTTPRequest.Header.Set("Content-Type", "application/vnd.api+json")
    return
  default:
    return

  }
}

// package private
func sendRequest(r *Request) (*http.Response, error) {
  r.Config.Logger.Debug("middleware: sendRequest executing")
  return r.Config.HTTPClient.Do(r.HTTPRequest)
}
