package tasks2


// IsLeapYear return true, if year is leap, and false if not
func IsLeapYear(year int) bool {
	return (year % 4 == 0 && year % 100 != 0) || (year % 100 == 0 && year % 400 == 0)
	
}
