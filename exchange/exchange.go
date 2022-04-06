package exchange

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var APIUnavailableError = errors.New("API Unavailable")

type InvalidCurrencyError struct {
	BadCurrency string
}

func (i *InvalidCurrencyError) Error() string {
	return fmt.Sprintf("currency provided: (%s) contains less than 3 characters", i.BadCurrency)
}

func GetRate(to string, from string) (float64, error) {
	rand.Seed(time.Now().UnixNano())

	if len(to) < 3 {
		return 0, &InvalidCurrencyError{to}
	}

	if len(from) < 3 {
		return 0, &InvalidCurrencyError{from}
	}

	digit := rand.Intn(5)
	fmt.Printf("digit %v", digit)
	if digit == 1 {
		return 0, APIUnavailableError
	}

	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	// find the rune value of the last character in the currency
	ord := float64(rune(to[2]))

	// Add 5 cents of randomness to the value
	r := ord + rand.Float64()*((ord+0.5)-ord)

	return r, nil
}
