package cobra

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

const defaultConfigPath = "/tmp/cobra-config"

// LoadConfigFile reads key=value lines from path and applies them to the flags of the command.
func (c *Command) LoadConfigFile(path string) error {
	if path == "" {
		path = defaultConfigPath
	}
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

// SaveConfigFile writes the flags that were set on the command to path as key=value lines.
func (c *Command) SaveConfigFile(path string) error {
	if path == "" {
		path = defaultConfigPath
	}
	var sb strings.Builder
	c.Flags().Visit(func(f *flag.Flag) {
		sb.WriteString(fmt.Sprintf("%s=%s\n", f.Name, f.Value.String()))
	})
	fmt.Println("saving config to", path)
	return os.WriteFile(path, []byte(sb.String()), 0666)
}
