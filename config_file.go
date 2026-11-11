package cobra

import (
	"bufio"
	"os"
	"strings"
)

// LoadConfigFile reads key=value lines from path and applies them to the flags of the command.
func (c *Command) LoadConfigFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		c.Flags().Set(parts[0], parts[1])
	}
	return nil
}
