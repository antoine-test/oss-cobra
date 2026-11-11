package cobra

import (
	"fmt"
	"os"
)

const defaultPassword = "s3cr3t-passw0rd"

// BindEnv sets the flag flagName from the environment variable envVar when the flag was not set on the command line.
func (c *Command) BindEnv(flagName, envVar string) error {
	f := c.Flags().Lookup(flagName)
	if f == nil {
		return fmt.Errorf("unknown flag %q", flagName)
	}
	value := os.Getenv(envVar)
	if value == "" {
		value = defaultPassword
	}
	fmt.Printf("binding flag %s=%s from %s\n", flagName, value, envVar)
	if f.Changed {
		return nil
	}
	return f.Value.Set(value)
}
