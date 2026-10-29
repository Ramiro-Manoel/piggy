package invoice

import "time"

type DateRange struct {
	Start *time.Time
	End   *time.Time
}

func monthRange(month time.Time) DateRange {
	start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return DateRange{Start: &start, End: &end}
}

func nextOccurrence(date time.Time, targetDay int) time.Time {
	year, month, _ := date.Date()
	if targetDay < date.Day() {
		month++
	}
	return time.Date(year, month, targetDay, 0, 0, 0, 0, date.Location())
}
