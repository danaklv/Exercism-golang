package tasks

import (
    "fmt"
    "errors"
) 

// TODO: define the 'DivideFood' function

func DivideFood(fd FodderCalculator, cows int) (float64, error) {
	amount, err := fd.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (amount*factor) / float64(cows), nil
   
}



// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(fd FodderCalculator, cows int) (float64, error) {
    if cows > 0 {
        return DivideFood(fd, cows)
    } 
    return 0, errors.New("invalid number of cows")
   
	}

// TODO: define the 'ValidateNumberOfCows' function

func ValidateNumberOfCows(cows int) (error) {
	if cows < 0 {
        str :=  fmt.Sprintf("%d cows are invalid: ", cows)
		var InvalidCowsError = errors.New(str + "there are no negative cows")
        return InvalidCowsError
	}
	if cows == 0 {
        str :=  fmt.Sprintf("%d cows are invalid: ", cows)
		var InvalidCowsError = errors.New(str + "no cows don't need food")
        return InvalidCowsError
	}
	return  nil

	
}

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}