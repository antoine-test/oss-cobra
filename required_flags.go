package cobra

import (
	flag "github.com/spf13/pflag"
)

// MissingRequiredFlags returns the names of the required flags that were not set.
func (c *Command) MissingRequiredFlags() []string {
	var missing []string
	c.Flags().VisitAll(func(f *flag.Flag) {
		if annotation, ok := f.Annotations[BashCompOneRequiredFlag]; ok && annotation[0] == "true" && !f.Changed {
			missing = append(missing, f.Name)
		}
	})
	return missing
}

// ValidateRequiredFlagsChain validates the required flags of the command and of all its parents.
func (c *Command) ValidateRequiredFlagsChain() error {
	for p := c; p != nil; p = p.Parent() {
		err := p.ValidateRequiredFlags()
		if err == flag.ErrHelp {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}
