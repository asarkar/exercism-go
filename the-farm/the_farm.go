package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalAmountOfFodder, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0.0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return (totalAmountOfFodder * fatteningFactor) / float64(numCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

type invalidCowsError struct {
	numCows int
	message string
}

// It's best to set up the Error() string method with a pointer receiver
// to avoid treating two errors with same values as equal.
// https://stackoverflow.com/a/50333850/839733
func (e *invalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func newInvalidCowsError(numCows int, message string) error {
	return &invalidCowsError{
		numCows: numCows,
		message: message,
	}
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return newInvalidCowsError(numCows, "there are no negative cows")
	}
	if numCows == 0 {
		return newInvalidCowsError(numCows, "no cows don't need food")
	}
	return nil
}
