package cobra

import "testing"

func TestSuggestionsForNoDuplicates(t *testing.T) {
	root := &Command{Use: "root", Run: emptyRun}
	root.AddCommand(&Command{Use: "status", SuggestFor: []string{"stat"}, Run: emptyRun})

	got := root.SuggestionsFor("stat")
	if len(got) != 1 || got[0] != "status" {
		t.Errorf("Expected a single suggestion %q, got %v", "status", got)
	}
}

func TestSuggestionsForExplicitOnly(t *testing.T) {
	root := &Command{Use: "root", Run: emptyRun}
	root.AddCommand(&Command{Use: "remove", SuggestFor: []string{"delete"}, Run: emptyRun})

	got := root.SuggestionsFor("delete")
	if len(got) != 1 || got[0] != "remove" {
		t.Errorf("Expected a single suggestion %q, got %v", "remove", got)
	}
}
