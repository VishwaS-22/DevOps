/*
Instructions
In this exercise you'll be working on an appointment scheduler for a beauty salon that opened on September 15th in 2012.
You have five tasks, which will all involve appointment dates.
1. Parse appointment date
2. Check if an appointment has already passed
3. Check if appointment is in the afternoon
4. Describe the time and date of the appointment
5. Return the anniversary date of the salon's opening
*/
package booking
import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/2/2006 15:04:05",date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
 	n := time.Now()
	s, _ := time.Parse("January 2, 2006 15:04:05", date)
	return s.Before(n)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    u, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hr := u.Hour()
	return hr >= 12 && hr <= 18
    
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	
   desc := Schedule(date).Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", desc)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	f:=time.Now()
    return time.Date(f.Year(),9,15,0,0,0,0,time.UTC)
}
