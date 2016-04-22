package cclog

import (
	"testing"
)


func TestCCLog(t *testing.T) {
	LogInstance.Init(LogLevel_Debug, true, true)
	Debug("Test", "Test Debug")
	Warning("Test", "Test Warning")
	Error("Test", "Test Error")
	Info("Test", "Test Info")
	Panic("Test", "Test Panic")
	Fatal("Test", "Test Fatal")
}
