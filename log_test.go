package gos

import (
	"testing"
)

func TestGolog(t *testing.T) {
	Debug("log debug test", "aa")
	Info("log info test")
	Warn("log warn test")
	Error("log error test")
	// Glog.Fatal("log fatal test")
}
