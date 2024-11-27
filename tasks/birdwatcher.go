package tasks

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
	for _, r := range birdsPerDay {
        count += r
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    index := 0
    sum := 0
    for week > 1  {
        index += 7
        week --
    }

    if len(birdsPerDay) - 6 >= 0 {
        for i := index;  i <= index + 6; i ++ {
        sum += birdsPerDay[i]
    }
    } else {
        for i := index;  i <= len(birdsPerDay) + 6; i ++ {
        sum += birdsPerDay[i]
    }
        }
  
    return sum
    
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, r := range birdsPerDay {
        if i == 0 || i % 2 == 0 {
           birdsPerDay[i] = r  + 1
        } 
    }
    return birdsPerDay
}