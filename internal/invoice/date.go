package invoice

import "time"

func nextOccurrence(date time.Time, targetDay int) time.Time {
	year, month, _ := date.Date()
	if targetDay < date.Day() {
		month++
	}
	return time.Date(year, month, targetDay, 0, 0, 0, 0, date.Location())
}
