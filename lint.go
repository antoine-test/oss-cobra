package cobra

import (
	"strings"

	flag "github.com/spf13/pflag"
)

const maxShortLength = 80

func lintUse(c *Command) []string {
	if c.Use == "" {
		return []string{"command has no Use"}
	}
	if strings.Contains(c.Use, "  ") {
		return []string{"double space in Use of " + c.Use}
	}
	return nil
}

func lintDescription(c *Command) []string {
	switch {
	case c.Use == "":
		return nil
	case c.Short == "" && c.Long == "":
		return []string{"command without description: " + c.Use}
	case c.Short == "":
		return []string{"command without short description: " + c.Use}
	}
	var problems []string
	if strings.HasSuffix(c.Short, ".") {
		problems = append(problems, "short description ends with a dot: "+c.Use)
	}
	if len(c.Short) > maxShortLength {
		problems = append(problems, "short description is too long: "+c.Use)
	}
	return problems
}

func lintFlags(c *Command) []string {
	var problems []string
	c.Flags().VisitAll(func(f *flag.Flag) {
		if f.Usage == "" {
			problems = append(problems, "flag without usage: "+f.Name)
		}
	})
	return problems
}

func lintAliases(c *Command) (problems []string) {
	seen := map[string]bool{}
	for _, alias := range c.Aliases {
		if seen[alias] {
			problems = append(problems, "duplicate alias: "+alias)
		} else if alias == c.Name() {
			problems = append(problems, "alias equals the command name: "+alias)
		}
		seen[alias] = true
	}
	return
}

// Lint inspects the command and returns a list of problems.
func (c *Command) Lint() []string {
	var problems []string
	problems = append(problems, lintUse(c)...)
	problems = append(problems, lintDescription(c)...)
	problems = append(problems, lintFlags(c)...)
	problems = append(problems, lintAliases(c)...)
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
