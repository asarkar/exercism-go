package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLogLineRegex   = regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)\]`)
	splitLogLineRegex   = regexp.MustCompile(`<[\^~*=-]*>`)
	quotedPasswordRegex = regexp.MustCompile(`(?i)\".*password.*\"`)
	endOfLineRegex      = regexp.MustCompile(`end-of-line\d+`)
	usernameRegex       = regexp.MustCompile(`User\s+([^\s]+)`)
)

func IsValidLine(text string) bool {
	return validLogLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	for _, line := range lines {
		if line != "" && quotedPasswordRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, 0)
	for _, line := range lines {
		if match := usernameRegex.FindStringSubmatch(line); match != nil {
			result = append(result, fmt.Sprintf("[USR] %s %s", match[1], line))
		} else {
			result = append(result, line)
		}
	}
	return result
}
