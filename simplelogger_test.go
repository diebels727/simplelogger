package simplelogger

import(
  "testing"
)

type MockWriter struct {
  LastMessage string
}

func (m *MockWriter) Write(p []byte) (n int,err error) {
  m.LastMessage = string(p)
  return len(p),nil
}

func TestInfo(t *testing.T) {
  m := &MockWriter{}
  logger := New(m,"[TEST] ",0)
  logger.Info("Info")
  if m.LastMessage != "[TEST] Info\n" {
    t.Error("Expected log message to be \"[TEST] Info\", got ",m.LastMessage)
  }
}

//
// func TestDebug(t *testing.T) {
//
//
// }