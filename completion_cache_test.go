package cobra

import (
	"testing"
	"time"
)

func TestCachedCompletionFunc(t *testing.T) {
	ClearCompletionCache()
	calls := 0
	f := CachedCompletionFunc(func(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective) {
		calls++
		return []Completion{"one", "two"}, ShellCompDirectiveNoFileComp
	}, time.Minute)
	cmd := &Command{Use: "cached"}

	f(cmd, nil, "o")
	got, directive := f(cmd, nil, "o")

	if calls != 1 {
		t.Errorf("Expected the wrapped function to be called once, got %d", calls)
	}
	if len(got) != 2 || directive != ShellCompDirectiveNoFileComp {
		t.Errorf("Unexpected result: %v, %v", got, directive)
	}
}

func TestCachedCompletionFuncExpires(t *testing.T) {
	ClearCompletionCache()
	calls := 0
	f := CachedCompletionFunc(func(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective) {
		calls++
		return nil, ShellCompDirectiveDefault
	}, time.Millisecond)
	cmd := &Command{Use: "cached"}

	f(cmd, nil, "")
	time.Sleep(10 * time.Millisecond)
	f(cmd, nil, "")

	if calls != 2 {
		t.Errorf("Expected the wrapped function to be called twice, got %d", calls)
	}
}

func TestRefreshCompletionsAsync(t *testing.T) {
	cmd := &Command{Use: "cached"}
	RefreshCompletionsAsync(func(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective) {
		return []Completion{"a"}, ShellCompDirectiveDefault
	}, cmd, nil, "", time.Minute)
	time.Sleep(50 * time.Millisecond)
}
