package thefarm

import "fmt"

// DivideFood calculates the amount of food per cow. Returns error if either method fails.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (totalFodder * factor) / float64(cows), nil
}

// ValidateInputAndDivideFood validates input before calling DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// InvalidCowsError is a custom error type for invalid cow numbers
type InvalidCowsError struct {
	Cows    int
	Message string
}

// Error implements the error interface
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.Cows, e.Message)
}

// ValidateNumberOfCows validates the number of cows and returns an error pointer if invalid
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			Cows:    cows,
			Message: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			Cows:    cows,
			Message: "no cows don't need food",
		}
	}
	return nil
}
