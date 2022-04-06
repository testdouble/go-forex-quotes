# Go Forex Quotes

Go library for fetching realtime forex quotes

### Install

```
go get github.com/testdouble/go-forex-quotes
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/testdouble/go-forex-quotes/client"
)

func main() {
	forexClient := client.New("API_KEY")

	quotes := forexClient.GetQuotes([]string{"EURUSD", "GBPJPY"})
	fmt.Printf("Quotes: %+v\n", quotes)

	usdToEur, err := forexClient.Convert("USD", "EUR", 100)
	fmt.Printf("$100 USD to EUR: %v. Err: %v\n", usdToEur, err)

	rate, err := forexClient.GetRate("JPY", "EUR")
	fmt.Printf("Rate for JPY to EUR: %v. Err: %v\n", rate, err)
}
```
