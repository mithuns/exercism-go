package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	dateFormat := "1/02/2006 15:04:05"
	schedule, _ := time.Parse(dateFormat, date)
	return schedule
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	dateFormat := "January 2, 2006 15:04:05"
	dateInQuestion, _ := time.Parse(dateFormat, date)
	return dateInQuestion.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	dateFormat := "Monday, January 2, 2006 15:04:05"
	dateInQuestion, _ := time.Parse(dateFormat, date)
	return dateInQuestion.Hour() >= 12 && dateInQuestion.Hour() <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dateFormat := "1/2/2006 15:04:05"
	dateInQuestion, _ := time.Parse(dateFormat, date)
	dateFormatOutput := "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.", dateInQuestion.Format(dateFormatOutput))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
