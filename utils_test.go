package utils

import (
	"errors"
	"testing"
)

func TestLoggers(t *testing.T) {
	LogInfo("info logger ran", "Type", "Function", "Block", "TestLoggers")
	LogError("error logger ran", errors.New("TestErrorString"), "Type", "Function", "Block", "TestLoggers")
	LogInfo("EmptyLogger")
	if !DidPanic(func() { LogInfo("info logger should not run", "bad args") }) {
		t.Fail()
	}
}
