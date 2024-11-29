package tasks2

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	patterns := []string{`^\[TRC\]`, `^\[DBG\]`, `^\[INF\]`, `^\[WRN\]`, `^\[ERR\]`, `^\[FTL\]`}
	for _, p := range patterns {
		r := regexp.MustCompile(p)
		if r.MatchString(text) {
			return true
		}
	}
	return false

}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))

	for i, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {

			username := matches[1]
			result[i] = fmt.Sprintf("[USR] %s %s", username, line)
		} else {
			result[i] = line
		}
	}

	return result
}
