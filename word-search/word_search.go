package wordsearch

import (
	"fmt"
)

// 8 possible search directions
//
//nolint:gochecknoglobals
var directions = [][2]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

// `Solve` finds all words in the puzzle, returning start and end coordinates.
// Although the problem description doesn't state clearly, a word may be searched
// only in one direction at a time.
//
// For example:
// Puzzle:
// +---+---+---+---+---+---+---+---+---+---+
// | j | e | f | b | l | p | e | p | r | e |
// +---+---+---+---+---+---+---+---+---+---+
// | t | c | l | o | j | u | r | e | r | m |
// +---+---+---+---+---+---+---+---+---+---+

// Word "clojure" horizontal match:
// +---+---+---+---+---+---+---+---+---+---+
// | j | e | f | b | l | p | e | p | r | e |
// +---+---+---+---+---+---+---+---+---+---+
// | t |[c]|[l]|[o]|[j]|[u]|[r]|[e]| r | m |
// +---+---+---+---+---+---+---+---+---+---+

// Alternative match that's not allowed because the
// direction changes from left-to-right to vertical:
// +---+---+---+---+---+---+---+---+---+---+
// | j | e | f | b | l | p |[e]| p | r | e |
// +---+---+---+---+---+---+---+---+---+---+
// | t |[c]|[l]|[o]|[j]|[u]|[r]| e | r | m |
// +---+---+---+---+---+---+---+---+---+---+
func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	results := make(map[string][2][2]int)
	m, n := len(puzzle), len(puzzle[0])

	for _, word := range words {
		found := false

		for r := 0; r < m && !found; r++ {
			for c := 0; c < n && !found; c++ {
				if puzzle[r][c] != word[0] {
					continue
				}
				if end := findWord(puzzle, word, r, c); end != nil {
					results[word] = [2][2]int{{c, r}, *end}
					found = true
				}
			}
		}
		if !found {
			return nil, fmt.Errorf("word %q not found", word)
		}
	}
	return results, nil
}

// `findWord` checks all directions from a given starting point.
func findWord(puzzle []string, word string, row, col int) *[2]int {
	for _, d := range directions {
		if end := follow(puzzle, word, row, col, d); end != nil {
			return end
		}
	}
	return nil
}

// `follow` follows a fixed direction until the word is matched.
func follow(puzzle []string, word string, r, c int, dir [2]int) *[2]int {
	m, n := len(puzzle), len(puzzle[0])
	for i := 1; i < len(word); i++ {
		r += dir[0]
		c += dir[1]
		if r < 0 || r >= m || c < 0 || c >= n || puzzle[r][c] != word[i] {
			return nil
		}
	}
	return &[2]int{c, r}
}
