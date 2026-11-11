package cobra

import "strings"

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
	return problems
}
