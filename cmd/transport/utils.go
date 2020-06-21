package transport

import (
  "fmt"
  "net/url"
)

func prepareUrl(base string, path string) *url.URL {
  URL, err := url.Parse(fmt.Sprintf("%s%s", base, path))
  if err != nil {
    panic(err)
  }
  return URL
}
