package transport

import (
  "testing"

  "github.com/stretchr/testify/assert"

  cfg "interview-accountapi/cmd"
)

func Test_ShouldValidateIfDataIsFilled(t *testing.T) {
  fixtures := []struct {
    input    *Request
    expected bool
  }{
    {
      input:    &Request{},
      expected: false,
    },
    {
      input:    &Request{Data: &MockStruct{Value: "test"}},
      expected: true,
    },
    {
      input:    &Request{Data: &struct {}{}},
      expected: true,
    },
  }
  for _, fixture := range fixtures {
    actual := fixture.input.DataFilled()
    assert.Equal(t, actual, fixture.expected)
  }
}

func Test_ShouldExecuteMiddleware(t *testing.T)  {
  var counter int
  expected := 2
  var MockBodyMiddleware  = NamedMiddleware{Name: "mock-body", Fn: func(r *Request) {
    counter++
  }}
  var PaginatorMiddleware  = NamedMiddleware{Name: "mock-body", Fn: func(r *Request) {
    counter++
  }}

  var mdw Middleware
  mdw.Body.Append(MockBodyMiddleware)
  mdw.Paginator.Append(PaginatorMiddleware)

  request := &Request{
    Config: cfg.Config{Logger: &MockLogger{}},
    Middleware: mdw,
  }

  err := request.Send()
  if err != nil {
    t.Fail()
  }

  assert.Equal(t, counter, expected)
}

func Test_ShouldWriteLoggerMessage(t *testing.T)  {
  var mdw Middleware
  var logger MockLogger
  expected := 3
  request := &Request{
    Config: cfg.Config{Logger: &logger},
    Middleware: mdw,
  }
  _ = request.Send()

  assert.Equal(t, logger.DebugExecuted, expected)
}
