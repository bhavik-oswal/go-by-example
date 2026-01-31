// Package expenses provides functions to track and analyze expenses.
package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns records that satisfy the predicate function.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

// ByDaysPeriod returns a predicate function for filtering by period.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns a predicate function for filtering by category.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns the total amount of expenses in a period.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	for _, r := range Filter(in, ByDaysPeriod(p)) {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total expenses in a category within a period.
// Returns an error if the category does not exist in any record.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	// Check if category exists at all
	categoryExists := false
	for _, r := range in {
		if r.Category == c {
			categoryExists = true
			break
		}
	}
	if !categoryExists {
		return 0, errors.New("unknown category " + c)
	}

	total := 0.0
	for _, r := range Filter(in, func(r Record) bool {
		return r.Category == c && r.Day >= p.From && r.Day <= p.To
	}) {
		total += r.Amount
	}
	return total, nil
}
