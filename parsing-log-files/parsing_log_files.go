package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	var validLogLines []*regexp.Regexp
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[TRC]`)))
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[DBG]`)))
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[INF]`)))
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[WRN]`)))
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[ERR]`)))
	validLogLines = append(validLogLines, regexp.MustCompile(regexp.QuoteMeta(`[FTL]`)))

	for i := range validLogLines {
		if indexes := validLogLines[i].FindIndex([]byte(text)); indexes != nil {
			if indexes[0] == 0 && indexes[1] == 5 {
				return true
			}
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for i := range lines {
		if regexp.MustCompile(`(?i).*".*password.*".*`).MatchString(lines[i]) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line[0-9]*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		s := regexp.MustCompile(`User +[^\s]+`).FindString(line)
		if s != "" {
			splits := strings.Split(s, " ")
			user := splits[len(splits)-1]
			lines[i] = fmt.Sprintf("[USR] %s %s", user, line)
		}
	}
	return lines
}
