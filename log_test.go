package log

import (
	"testing"
	"time"
)


func TestLog(t *testing.T) {
	Init(&Config{level: DEBUG})
	Debug("Test", "Test Debug")
	Warn("Test", "Test Trace")
	Error("Test", "Test Error")
	Info("Test", "Test Info")
	Fatal("Test", "Test Fatal")
	time.Sleep(3*time.Second)
}
