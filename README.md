# Go Forex Quotes

Go library for fetching realtime forex quotes

### Install

```
go get github.com/testdouble/go-forex-quotes
```

### Usage

```
package main

import (
  "fmt"
  "github.com/testdouble/go-forex-quotes"
)

func main() {
  client := forex.New("API_KEY")

  quotes := client.GetQuotes(["EURUSD", "GBPJPY"])
  usdToEur := client.Convert("USD", "EUR", 100)

  fmt.Printf("Quotes: %v\n", quotes)
  fmt.Printf("$100 USD to EUR: %v\n", usdToEur)
}
```
