package robotname

import (
	"errors"
	"math/rand/v2"
	"strings"
)

//nolint:gochecknoglobals
var (
	names = make(map[string]struct{})
	n     = 26 * 26 * 10 * 10 * 10
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(names) == n {
		return "", errors.New("namespace exhausted")
	}

	var sb strings.Builder
	for range 2 {
		sb.WriteRune(rune('A' + rand.IntN(26)))
	}
	for range 3 {
		sb.WriteByte(byte('0' + rand.IntN(10)))
	}

	name := sb.String()
	if _, exists := names[name]; exists {
		return r.Name() // collision, retry
	}

	names[name] = struct{}{}
	r.name = name
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	_, _ = r.Name()
}
