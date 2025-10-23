package meetup

import (
	"fmt"
	"time"
)

// Define the WeekSchedule type here.
//
// Go does not have a built-in enum keyword like some other languages.  Instead, we can emulate
// enum-like behavior using a combination of custom types and constants, often with the help of
// `iota`.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// addDays returns a closure function that, when called with an integer `days`,
// returns a new `time.Time` value offset by that many days from the original time `t`.
//
// Example:
//
//	base := time.Now() 	 // 2025-10-22 00:00:00 +0000 UTC
//	add := addDays(base)
//	fmt.Println(add(3))  // 2025-10-25 00:00:00 +0000 UTC
//	fmt.Println(add(-2)) // 2025-10-20 00:00:00 +0000 UTC
func addDays(t time.Time) func(int) time.Time {
	return func(days int) time.Time {
		return t.AddDate(0, 0, days)
	}
}

// firstWeekday returns the date of the first occurrence of the specified `weekday`
// in a given `month` and `year`.
//
// Example:
//
//	first := firstWeekday(time.Friday, time.November, 2025)
//	fmt.Println(first) // 2025-11-07 00:00:00 +0000 UTC
func firstWeekday(weekday time.Weekday, month time.Month, year int) time.Time {
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	dayOfWeek := firstOfMonth.Weekday()
	offset := weekday - dayOfWeek
	if offset < 0 {
		offset += 7
	}
	return firstOfMonth.AddDate(0, 0, int(offset))
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	weekday := firstWeekday(wDay, month, year)
	addDays := addDays(weekday)
	switch wSched {
	case First:
		return addDays(0).Day()
	case Second:
		return addDays(7).Day()
	case Third:
		return addDays(14).Day()
	case Fourth:
		return addDays(21).Day()
	case Last:
		lastWeekday := addDays(28)
		if lastWeekday.Month() == month {
			return lastWeekday.Day()
		}
		return addDays(21).Day()
	case Teenth:
		for i := 7; weekday.Day() < 13; i += 7 {
			weekday = addDays(i)
		}
		return weekday.Day()
	default:
		panic(fmt.Sprintf("invalid WeekSchedule: %v", wSched))
	}
}
