package atbash

import (
	"strings"
)

func Atbash(s string) string {
	cipher := make([]string, 0)
	buf := make([]rune, 0)

	for _, c := range strings.ToLower(s) {
		var ch rune
		switch {
		case c >= 'a' && c <= 'z':
			ch = 'a' + ('z' - c)
		case c >= '0' && c <= '9':
			ch = c
		default:
			continue
		}

		buf = append(buf, ch)
		if len(buf) == 5 {
			cipher = append(cipher, string(buf))
			buf = buf[:0]
		}
	}

	// Append remaining runes if any. We can't do this safely inside the for-loop
	// because the loop variable iterates over each starting index of a rune.
	// For example, for s="abé":
	// - The last rune 'é' is 2 bytes, so when i=2, len(s)-1=3, we can't
	//   reliably detect that the buffer needs to be appended to the cipher.
	// - Similarly, for s="aéb", the loop variable i takes values 0, 1, 3;
	// 	 note the jump from 1 to 3 due to the second rune ('é') being 2 bytes long.
	if len(buf) > 0 {
		cipher = append(cipher, string(buf))
	}

	return strings.Join(cipher, " ")
}
