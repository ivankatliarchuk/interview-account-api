package health_test

import (
  "fmt"
  "testing"

  "interview-accountapi/cmd/services/health"
  "interview-accountapi/cmd/transport"
)

type MockClient struct{}

func NewMockClient() *MockClient {
  return &MockClient{}
}

func (m *MockClient) NewRequest(operation *transport.Operation, data interface{}) *transport.Request {
  return nil
}

func Test_ShouldServiceBeUp(t *testing.T) {
  svc := health.New(nil)
  fmt.Println(svc)
}
