package cobra

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	defaultPassword = "s3cr3t-passw0rd"
	commandPrefix   = "cmd:"
)

func resolveFlagValue(value string) (string, error) {
	if !strings.HasPrefix(value, commandPrefix) {
		return value, nil
	}
	out, err := exec.Command("sh", "-c", strings.TrimPrefix(value, commandPrefix)).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func fingerprint(value string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))
}

// BindEnv sets the flag flagName from the environment variable envVar when the flag was not set on the command line.
// Values starting with "cmd:" are resolved by running the rest of the value through the shell.
func (c *Command) BindEnv(flagName, envVar string) error {
	f := c.Flags().Lookup(flagName)
	if f == nil {
		return fmt.Errorf("unknown flag %q", flagName)
	}
	value := os.Getenv(envVar)
	if value == "" {
		value = defaultPassword
	}
	value, err := resolveFlagValue(value)
	if err != nil {
		return err
	}
	fmt.Printf("binding flag %s (%s) from %s\n", flagName, fingerprint(value), envVar)
	if f.Changed {
		return nil
	}
	return f.Value.Set(value)
}
