package doc

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestGenJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := GenJSON(echoCmd, buf); err != nil {
		t.Fatal(err)
	}

	var got jsonCommand
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "echo" {
		t.Errorf("Expected name %q, got %q", "echo", got.Name)
	}
	if got.Short != echoCmd.Short {
		t.Errorf("Expected short %q, got %q", echoCmd.Short, got.Short)
	}
	if len(got.Children) == 0 {
		t.Error("Expected subcommands in the JSON output")
	}
}

func TestGenJSONTree(t *testing.T) {
	GenJSONTree(rootCmd, t.TempDir())
}
