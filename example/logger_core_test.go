package example

import (
	logger "logger-go"
	"testing"
)

func TestCoreLogger(t *testing.T) {
	err := logger.Init().SetLevel("debug").SetServer("cannon").Call()
	if err != nil {
		t.Fatal(err)
		return
	}
	logger.Debug("hello cannon")

	logger.WithStr("uuid", "===cannon 123===")
	logger.Debugf("hello %s", "cannon")
}
