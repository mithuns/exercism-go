package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredResults := make([]Record, 0)
	for _, e := range in {
		if predicate(e) {
			filteredResults = append(filteredResults, e)
		}
	}
	return filteredResults
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return p.From <= record.Day && p.To >= record.Day
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	results := Filter(in, ByDaysPeriod(p))
	for index := range results {
		total += results[index].Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	total := 0.0
	results := Filter(in, ByCategory(c))
	if len(results) == 0 {
		return 0.0, errors.New("No records in the list")
	}
	results = Filter(results, ByDaysPeriod(p))

	for index := range results {
		total += results[index].Amount
	}
	return total, nil
}
