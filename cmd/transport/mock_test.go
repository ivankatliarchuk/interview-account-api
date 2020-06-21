package transport

type MockStruct struct {
  Value string `json:"value"`
}

type MockLogger struct {
  DebugExecuted int
}

func (m *MockLogger) Info(...interface{}) {}
func (m *MockLogger) Debug(...interface{}) {
  m.DebugExecuted++
}
func (m *MockLogger) Error(...interface{}) {}

