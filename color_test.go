package clog

import (
	"testing"
)

func TestColor(t *testing.T) {
	Info("info message")
	Infof("%s", "info message")
	Warn("warn message")
	Warnf("%s", "warn message")
	Error("error message")
	Errorf("%s", "error message")
}
