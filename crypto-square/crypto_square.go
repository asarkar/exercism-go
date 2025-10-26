package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	text := normalize(pt)
	width := int(math.Ceil(math.Sqrt(float64(len(text)))))
	chunks := chunk(text, width)
	secret := zip(chunks, width)
	return strings.Join(secret, " ")
}

func normalize(text string) string {
	var sb strings.Builder
	for _, c := range text {
		if isAlphaNum(c) {
			sb.WriteRune(unicode.ToLower(c))
		}
	}
	return sb.String()
}

func isAlphaNum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func chunk(text string, width int) []string {
	n := len(text)
	chunks := make([]string, 0)
	for i := 0; i < n; i += width {
		chunks = append(chunks, text[i:min(i+width, n)])
	}
	return chunks
}

func zip(chunks []string, width int) []string {
	var sb strings.Builder
	secret := make([]string, 0)
	for col := range width {
		for _, chunk := range chunks {
			if col < len(chunk) {
				sb.WriteByte(chunk[col])
			} else {
				sb.WriteByte(' ')
			}
		}
		secret = append(secret, sb.String())
		sb.Reset()
	}
	return secret
}
