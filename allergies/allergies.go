package allergies

import "slices"

func allergens() []string {
	return []string{
		"eggs", "peanuts", "shellfish", "strawberries",
		"tomatoes", "chocolate", "pollen", "cats",
	}
}

func Allergies(allergies uint) []string {
	a := allergens()
	allergies %= 256
	// Go >= 1.23.
	// https://stackoverflow.com/a/78185810/839733
	// https://bitfieldconsulting.com/posts/iterators
	return slices.Collect(
		func(yield func(string) bool) {
			for i := 0; allergies > 0; i++ {
				if allergies&1 == 1 {
					if !yield(a[i]) {
						return
					}
				}
				allergies >>= 1
			}
		},
	)
}

func AllergicTo(allergies uint, allergen string) bool {
	return slices.Contains(Allergies(allergies), allergen)
}
