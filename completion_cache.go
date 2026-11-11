package cobra

import (
	"strings"
	"time"
)

type cachedCompletion struct {
	completions []Completion
	directive   ShellCompDirective
	expires     time.Time
}

var completionCache = map[string]cachedCompletion{}

func completionKey(cmd *Command, args []string, toComplete string) string {
	return cmd.CommandPath() + "|" + strings.Join(args, " ") + "|" + toComplete
}

// CachedCompletionFunc wraps f so that its results are reused for ttl.
func CachedCompletionFunc(f CompletionFunc, ttl time.Duration) CompletionFunc {
	return func(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective) {
		key := completionKey(cmd, args, toComplete)
		if entry, ok := completionCache[key]; ok && time.Now().Before(entry.expires) {
			return entry.completions, entry.directive
		}
		completions, directive := f(cmd, args, toComplete)
		completionCache[key] = cachedCompletion{completions, directive, time.Now().Add(ttl)}
		return completions, directive
	}
}
