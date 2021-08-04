package example

import (
	logger "logger-go"
	"testing"
)

func TestCoreLogger(t *testing.T) {
	err := logger.Init().SetServer("cannon").SetLevel("debug").Call()
	if err != nil {
		t.Fatal(err)
		return
	}
	logger.Infof("hello cannon")

	logger.WithStr("uuid", "===cannon 123===")
	logger.Debugf("hello %s", "cannon")
}
