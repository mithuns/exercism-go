package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var recommendation rune = '❗'
var search rune = '🔍'
var weather rune = '☀'

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendationIndex := strings.IndexRune(log, recommendation)
	searchIndex := strings.IndexRune(log, search)
	weatherIndex := strings.IndexRune(log, weather)

	if recommendationIndex >= 0 && searchIndex >= 0 {
		if recommendationIndex == 0 {
			return "recommendation"
		}
		if searchIndex == 0 {
			return "search"
		}
	}
	if recommendationIndex >= 0 && weatherIndex >= 0 {
		if recommendationIndex == 0 {
			return "recommendation"
		}
		if weatherIndex >= 0 {
			return "weather"
		}
	}
	if weatherIndex >= 0 && searchIndex >= 0 {
		if weatherIndex == 0 {
			return "weather"
		}
		if searchIndex >= 0 {
			return "search"
		}
	}
	if recommendationIndex >= 0 {
		return "recommendation"
	}
	if searchIndex >= 0 {
		return "search"
	}
	if weatherIndex >= 0 {
		return "weather"
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
