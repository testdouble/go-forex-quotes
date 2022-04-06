package client

import (
	"github.com/testdouble/go-forex-quotes/exchange"
)

type Client struct {
	apiKey string
}

type Quote struct {
	Currencies string
	Rate       float64
	Error      error
}

func New(apiKey string) *Client {
	return &Client{
		apiKey,
	}
}

func (c *Client) GetRate(to string, from string) (float64, error) {
	if to == from {
		return 1.0, nil
	}

	return exchange.GetRate(to, from)
}

func (c *Client) GetQuotes(currencies []string) (quotes []Quote) {
	for _, cs := range currencies {
		var rate float64
		to := cs[0:3]
		from := cs[3:6]
		rate, err := exchange.GetRate(to, from)
		quotes = append(quotes, Quote{Currencies: cs, Rate: rate, Error: err})
	}
	return
}

func (c *Client) Convert(to string, from string, amount float64) (float64, error) {
	rate, err := exchange.GetRate(to, from)
	total := amount * rate

	return total, err
}
