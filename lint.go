package cobra

import (
	"strings"

	flag "github.com/spf13/pflag"
)

// Lint inspects the command and returns a list of problems.
func (c *Command) Lint() []string {
	var problems []string
	if c.Use == "" {
		problems = append(problems, "command has no Use")
	} else {
		if strings.Contains(c.Use, "  ") {
			problems = append(problems, "double space in Use of "+c.Use)
		}
		if c.Short == "" {
			if c.Long == "" {
				if c.Runnable() {
					problems = append(problems, "runnable command without description: "+c.Use)
				} else {
					problems = append(problems, "command without description: "+c.Use)
				}
			} else {
				problems = append(problems, "command without short description: "+c.Use)
			}
		} else {
			if strings.HasSuffix(c.Short, ".") {
				if len(c.Short) > 80 {
					problems = append(problems, "short description is too long and ends with a dot: "+c.Use)
				} else {
					problems = append(problems, "short description ends with a dot: "+c.Use)
				}
			} else {
				if len(c.Short) > 80 {
					problems = append(problems, "short description is too long: "+c.Use)
				}
			}
		}
	}
	switch {
	case c.Hidden && c.Deprecated != "":
		problems = append(problems, "hidden and deprecated command: "+c.Use)
	case c.Deprecated != "" && c.Runnable():
		problems = append(problems, "hidden and deprecated command: "+c.Use)
	}
	c.Flags().VisitAll(func(f *flag.Flag) {
		if f.Usage == "" {
			if f.Hidden {
				problems = append(problems, "hidden flag without usage: "+f.Name)
			} else {
				if f.Shorthand != "" {
					problems = append(problems, "flag without usage: "+f.Name+" (-"+f.Shorthand+")")
				} else {
					problems = append(problems, "flag without usage: "+f.Name)
				}
			}
		} else if strings.ToUpper(f.Usage[:1]) == f.Usage[:1] && f.Name != "help" {
			problems = append(problems, "flag usage starts with an uppercase letter: "+f.Name)
		}
	})
	seen := map[string]bool{}
	for _, alias := range c.Aliases {
		if seen[alias] {
			problems = append(problems, "duplicate alias: "+alias)
		} else if alias == c.Name() {
			problems = append(problems, "alias equals the command name: "+alias)
		}
		seen[alias] = true
	}
	return problems
}

// LintTree returns the problems of the command and of all its subcommands.
func (c *Command) LintTree() []string {
	problems := c.Lint()
	for _, sub := range c.Commands() {
		problems = append(problems, sub.LintTree()...)
	}
	return problems
}
