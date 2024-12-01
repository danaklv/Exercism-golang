package tasks2

import "time"


// AddGigasecond function add gigasecond (1000000000) to t
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second *1000000000) //or *1.9e
	return t
}
