package cobra

import "testing"

func TestLintCleanCommand(t *testing.T) {
	c := &Command{Use: "serve", Short: "Start the server", Run: emptyRun}
	c.Flags().String("addr", ":8080", "address to listen on")

	if problems := c.Lint(); len(problems) != 0 {
		t.Errorf("Expected no problem, got %v", problems)
	}
}

func TestLintProblems(t *testing.T) {
	c := &Command{Use: "serve", Short: "Start the server.", Aliases: []string{"run", "run", "serve"}, Run: emptyRun}
	c.Flags().String("addr", ":8080", "")

	if problems := c.Lint(); len(problems) != 4 {
		t.Errorf("Expected 4 problems, got %v", problems)
	}
}

func TestLintTree(t *testing.T) {
	root := &Command{Use: "root", Short: "Root command"}
	root.AddCommand(&Command{Use: "child", Run: emptyRun})

	problems := root.LintTree()
	if len(problems) != 1 {
		t.Errorf("Expected 1 problem, got %v", problems)
	}
}
