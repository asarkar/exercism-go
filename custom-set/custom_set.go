package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

//nolint:recvcheck // This exercises mixes receiver and non-pointer receivers
type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(list []string) Set {
	s := New()
	for i := range list {
		s.Add(list[i])
	}
	return s
}

func (s Set) String() string {
	quoted := make([]string, len(s))
	for i := range s {
		quoted[i] = fmt.Sprintf("%q", s[i]) // adds quotes
	}
	return fmt.Sprintf("{%s}", strings.Join(quoted, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for i := range s {
		if s[i] == elem {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		*s = append(*s, elem)
	}
}

// Subset returns true if s1 is a subset of s2.
func Subset(s1, s2 Set) bool {
	for i := range s1 {
		if !s2.Has(s1[i]) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common.
func Disjoint(s1, s2 Set) bool {
	for i := range s1 {
		if s2.Has(s1[i]) {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 contain the same elements.
func Equal(s1, s2 Set) bool {
	return &s1 == &s2 || (len(s1) == len(s2) && Subset(s1, s2))
}

// Intersection returns a new set with elements common to s1 and s2.
func Intersection(s1, s2 Set) Set {
	res := New()
	for i := range s1 {
		if s2.Has(s1[i]) {
			res.Add(s1[i])
		}
	}
	return res
}

// Difference returns a new set with elements in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	res := New()
	for i := range s1 {
		if !s2.Has(s1[i]) {
			res.Add(s1[i])
		}
	}
	return res
}

func Union(s1, s2 Set) Set {
	// Make a copy so as not to affect `s1`.
	res := NewFromSlice(s1)
	for _, e := range s2 {
		res.Add(e)
	}
	return res
}
