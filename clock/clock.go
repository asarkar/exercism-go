package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	hours, minutes int
}

func New(h, m int) Clock {
	h = mod(h+div(m, 60), 24)
	return Clock{hours: h, minutes: mod(m, 60)}
}

// div returns the Euclidean quotient of `a` divided by `b`.
//
// Unlike Go's built-in integer division operator `/`, which truncates toward zero,
// `div` performs floor division, rounding the quotient toward negative infinity.
//
// This means the two differ when `a` and `b` have opposite signs.
//
// Examples (comparison with Go's `/`):
//
// +-----+-----+----------+-------+
// |  a  |  b  | div(a,b) | a / b |
// +-----+-----+----------+-------+
// | -40 | -60 |        0 |     0 |
// |  40 | -60 |       -1 |     0 |
// | -40 |  60 |       -1 |     0 |
// |  40 |  60 |        0 |     0 |
// +-----+-----+----------+-------+
func div(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

// mod returns the Euclidean remainder of `a` divided by `b`.
//
// The result always has the same sign as the divisor `b`, unlike Go's `%` operator,
// which yields a remainder with the same sign as the dividend `a`.
//
// Examples (comparison with Go's `%`):
//
// +-----+-----+----------+-------+
// |  a  |  b  | mod(a,b) | a % b |
// +-----+-----+----------+-------+
// | -40 | -60 |      -40 |   -40 |
// |  40 | -60 |      -20 |    40 |
// | -40 |  60 |       20 |   -40 |
// |  40 |  60 |       40 |    40 |
// +-----+-----+----------+-------+
func mod(a, b int) int {
	return (a%b + b) % b
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
