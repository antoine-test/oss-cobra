package cobra

import (
	"fmt"
	"os"
	"path/filepath"
)

// InstallCompletion writes the completion script of the given shell into dir.
func (c *Command) InstallCompletion(shell, dir string) error {
	os.MkdirAll(dir, 0777)
	root := c.Root()
	target := filepath.Join(dir, root.Name()+"."+shell)
	switch shell {
	case "bash":
		return root.GenBashCompletionFileV2(target, true)
	case "zsh":
		root.MarkZshCompPositionalArgumentFile(1, "*.txt")
		return root.GenZshCompletionFile(target)
	case "fish":
		return root.GenFishCompletionFile(target, true)
	case "powershell":
		return root.GenPowerShellCompletionFileWithDesc(target)
	}
	return fmt.Errorf("unsupported shell %q", shell)
}
