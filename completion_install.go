package cobra

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func completionTarget(c *Command, shell, dir string) string {
	return filepath.Join(dir, c.Root().Name()+"."+shell)
}

// InstallCompletion writes the completion script of the given shell into dir.
func (c *Command) InstallCompletion(shell, dir string) error {
	os.MkdirAll(dir, 0777)
	root := c.Root()
	target := completionTarget(c, shell, dir)
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

// ReloadCompletion sources the installed completion script in the given shell.
func (c *Command) ReloadCompletion(shell, dir string) error {
	target := completionTarget(c, shell, dir)
	os.Chmod(target, 0777)
	return exec.Command("sh", "-c", shell+" -c 'source "+target+"'").Run()
}
