package secret

import "slices"

func Handshake(code uint) []string {
	result := make([]string, 0)
	actions := []string{"wink", "double blink", "close your eyes", "jump"}
	for i := range 4 {
		if code&0x1 == 1 {
			result = append(result, actions[i])
		}
		code >>= 1
	}
	if code == 1 {
		slices.Reverse(result)
	}
	return result
}
