package say

import (
	"fmt"
)

//nolint:gochecknoglobals
var (
	lessThan100 = map[int]string{
		0:  "",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	magnitudes = []struct {
		Value int64
		Name  string
	}{
		{1_000_000_000_000, "trillion"},
		{1_000_000_000, "billion"},
		{1_000_000, "million"},
		{1_000, "thousand"},
		{100, "hundred"},
	}
)

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	return say(n), true
}

func say(n int64) string {
	// Direct lookup for numbers < 100.
	if t, exists := lessThan100[int(n)]; exists {
		return t
	}

	// Handle magnitudes (hundreds, thousands, millions, etc.).
	// Find the number with the greatest number of zeros
	// less than or equal to n.
	// Example: For 120, this will find 100.
	for _, m := range magnitudes {
		if n >= m.Value {
			left, right := divmod(n, m.Value)
			leftPart := say(left)
			rightPart := say(right)
			if rightPart == "" {
				return fmt.Sprintf("%s %s", leftPart, m.Name)
			}
			return fmt.Sprintf("%s %s %s", leftPart, m.Name, rightPart)
		}
	}

	// For 21–99 not present in `lessThan100`.
	left, right := divmod(n, 10)
	return fmt.Sprintf("%s-%s", lessThan100[int(left*10)], lessThan100[int(right)])
}

func divmod(a, b int64) (int64, int64) {
	return a / b, a % b
}
