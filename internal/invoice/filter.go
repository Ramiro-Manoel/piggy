package invoice

type Filter struct {
	CardIDs   []string
	CloseDate DateRange
	PayDate   DateRange
}
