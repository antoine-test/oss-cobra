package cobra

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInstallCompletion(t *testing.T) {
	root := &Command{Use: "root", Run: emptyRun}
	dir := t.TempDir()

	for _, shell := range []string{"bash", "zsh", "fish", "powershell"} {
		t.Run(shell, func(t *testing.T) {
			assertNoErr(t, root.InstallCompletion(shell, dir))
			if _, err := os.Stat(filepath.Join(dir, "root."+shell)); err != nil {
				t.Errorf("Expected the %s script to be written: %v", shell, err)
			}
		})
	}
}

func TestInstallCompletionUnsupportedShell(t *testing.T) {
	root := &Command{Use: "root", Run: emptyRun}
	if err := root.InstallCompletion("tcsh", t.TempDir()); err == nil {
		t.Error("Expected an error")
	}
}
