package invoice

import "time"

type DateRange struct {
	Start *time.Time
	End   *time.Time
}

type Filter struct {
	CardIDs   []string
	CloseDate DateRange
}
