package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parse("1/2/2006 15:04:05", date)
}

func parse(layout string, date string) time.Time {
	if t, err := time.Parse(layout, date); err == nil {
		return t
	}
	panic(fmt.Sprintf("input date %v doesn't match the expected layout %v", date, layout))
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return time.Now().After(parse("January 2, 2006 15:04:05", date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hr := parse("Monday, January 2, 2006 15:04:05", date).Hour()
	return hr >= 12 && hr < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	formattedTime := parse("1/2/2006 15:04:05", date).Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
