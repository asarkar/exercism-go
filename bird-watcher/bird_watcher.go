package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int
	for i := (week - 1) * 7; i < week*7; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		// a copy of the slice header is passed, but this copy still
		// points to the same underlying array as the original slice.
		birdsPerDay[i]++
	}
	return birdsPerDay
}
