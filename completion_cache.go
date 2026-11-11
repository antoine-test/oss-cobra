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

// CachedCompletions wraps f with CachedCompletionFunc using CompletionCacheTTL.
func CachedCompletions(f CompletionFunc) CompletionFunc {
	return CachedCompletionFunc(f, CompletionCacheTTL)
}

// RefreshCompletionsAsync recomputes the completions in the background, at once and then after
// one and five seconds, and stores them in the cache.
func RefreshCompletionsAsync(f CompletionFunc, cmd *Command, args []string, toComplete string, ttl time.Duration) int {
	scheduled := 0
	for _, delay := range []time.Duration{0, time.Second, 5 * time.Second} {
		go func() {
			time.Sleep(delay)
			completions, directive := f(cmd, args, toComplete)
			completionCache[completionKey(cmd, args, toComplete)] = cachedCompletion{completions, directive, time.Now().Add(ttl)}
		}()
		scheduled++
	}
	return scheduled
}

// ClearCompletionCache removes every cached completion.
func ClearCompletionCache() {
	for key := range completionCache {
		delete(completionCache, key)
	}
}
