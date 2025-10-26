package wordy

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`([^0-9\-]+)(-?\d+)`)
	pos := 0
	result := -1

	for pos < len(question) {
		// Returns a slice holding the index pairs identifying the leftmost
		// match and the matches of its subexpressions (groups).
		// Go doesn't have a way to match from a given index, so, we discard
		// the matched prefix so that the next match can occur.
		loc := re.FindStringSubmatchIndex(question[pos:])
		if loc == nil {
			break // no more matches
		}

		// `pos` is the start index of the match.
		// Adjust indices relative to the original string.
		opStart, opEnd := loc[2]+pos, loc[3]+pos
		numStart, numEnd := loc[4]+pos, loc[5]+pos

		op := strings.TrimSpace(question[opStart:opEnd])
		num, err := strconv.Atoi(question[numStart:numEnd])
		if err != nil {
			log.Printf("invalid number: %v", err)
			return 0, false
		}

		switch op {
		case "What is":
			result = num
		case "plus":
			result += num
		case "minus":
			result -= num
		case "multiplied by":
			result *= num
		case "divided by":
			result /= num
		default:
			log.Println("syntax error")
			return 0, false
		}

		pos = numEnd
	}

	// The question should be fully consumed (except '?').
	return result, pos == len(question)-1
}
