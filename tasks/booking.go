package tasks

import (
    "time"
    "fmt"
    )

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout,date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout,date)
	today := time.Now()  
    return t.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, _ := time.Parse(layout,date)
    return t.Hour() >= 12 && t.Hour() < 18
    
}
// Description("7/25/2019 13:45:00")
// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout,date)
    
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", 
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
   
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().UTC().Year()
	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}