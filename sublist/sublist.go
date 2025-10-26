package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	n := min(len(l1), len(l2))
	lps := make([]int, n)

	if len(l2) == n {
		kmpLoop(l2, l2, 1, lps, true)
		i := kmpLoop(l1, l2, 0, lps, false)
		if i == n && len(l1) == n {
			return RelationEqual
		}
		if i == n {
			return RelationSuperlist
		}
		return RelationUnequal
	}
	rel := Sublist(l2, l1)
	if rel == RelationSuperlist {
		return RelationSublist
	}
	return rel
}

func kmpLoop(haystack, needle []int, start int, lps []int, buildLPS bool) int {
	i, j := 0, start
	for i < len(needle) && j < len(haystack) {
		switch {
		case needle[i] == haystack[j]:
			i++
			j++
			if buildLPS {
				lps[i] = i
			}
		case i > 0:
			i = lps[i-1]
		default:
			if buildLPS {
				lps[i] = 0
			}
			j++
		}
	}
	return i
}
