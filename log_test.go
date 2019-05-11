package log

import (
	"testing"
)


func TestLog(t *testing.T) {
	Debug("Test", "Test Debug")
	Error("Test", "Test Error")
	Info("Test", "Test Info")
	Trace("Test", "Test Trace")
	Fatal("Test", "Test Fatal")
}
