package cobra

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config")
	if err := os.WriteFile(path, []byte("# comment\nname=cobra\n\ncount=3\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")
	c.Flags().Int("count", 0, "")

	if err := c.LoadConfigFile(path); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if got, _ := c.Flags().GetString("name"); got != "cobra" {
		t.Errorf("Expected name %q, got %q", "cobra", got)
	}
	if got, _ := c.Flags().GetInt("count"); got != 3 {
		t.Errorf("Expected count 3, got %d", got)
	}
}

func TestLoadConfigFileInvalidLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config")
	if err := os.WriteFile(path, []byte("name\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")

	if err := c.LoadConfigFile(path); err == nil {
		t.Fatal("Expected an error")
	}
}

func TestSaveConfigFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config")
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")
	assertNoErr(t, c.Flags().Set("name", "cobra"))

	assertNoErr(t, c.SaveConfigFile(path))
}
