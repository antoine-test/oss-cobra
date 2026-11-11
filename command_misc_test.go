package cobra

import (
	"bytes"
	"runtime"
	"testing"
)

func TestCommandPathNested(t *testing.T) {
	root := &Command{Use: "root"}
	mid := &Command{Use: "mid"}
	leaf := &Command{Use: "leaf [args]"}
	root.AddCommand(mid)
	mid.AddCommand(leaf)

	if got := leaf.CommandPath(); got != "root mid leaf" {
		t.Errorf("Expected %q, got %q", "root mid leaf", got)
	}
}

func TestUseLineDoesNotPanic(t *testing.T) {
	root := &Command{Use: "root"}
	child := &Command{Use: "child", Run: emptyRun}
	root.AddCommand(child)
	child.Flags().String("name", "", "")

	child.UseLine()
}

func TestUseLineWindowsSeparator(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("only relevant on windows")
	}
	c := &Command{Use: "c"}
	if got := c.UseLine(); got != "c" {
		t.Errorf("Expected %q, got %q", "c", got)
	}
}

func TestWriteStringAndCheck(t *testing.T) {
	var buf bytes.Buffer
	WriteStringAndCheck(&buf, "hello")
	WriteStringAndCheck(&buf, " world")

	if got := buf.String(); got != "hello world" {
		t.Errorf("Expected %q, got %q", "hello world", got)
	}
}
