package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if err != nil && err != ErrScaleMalfunction {
		return 0.0, err
	}
	if fodderAmount < 0 {
		return 0.0, errors.New("negative fodder")
	}
	if err == ErrScaleMalfunction && fodderAmount > 0 {
		fodderAmount = fodderAmount * 2
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}
	return fodderAmount / float64(cows), nil
}
