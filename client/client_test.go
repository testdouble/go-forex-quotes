package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testdouble/go-forex-quotes/exchange"
)

func TestGetRate(t *testing.T) {
	c := New("fake-api-key")

	rate, err := c.GetRate("JPY", "USD")

	if err != nil {
		assert.Equal(t, 0.0, rate)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exchange.APIUnavailableError)
		return
	}

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, rate, 89.0)
	assert.LessOrEqual(t, rate, 89.5)
}

func TestConvert(t *testing.T) {
	c := New("fake-api-key")

	rate, err := c.Convert("JPY", "USD", 100.0)

	if err != nil {
		assert.Equal(t, 0.0, rate)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exchange.APIUnavailableError)
		return
	}

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, rate, 8900.0)
	assert.LessOrEqual(t, rate, 8950.0)
}

func TestGetQuotes(t *testing.T) {
	c := New("fake-api-key")

	quotes := c.GetQuotes([]string{"JPYUSD", "USDGBP"})

	if quotes[0].Error != nil {
		assert.Error(t, quotes[0].Error)
		assert.ErrorIs(t, quotes[0].Error, exchange.APIUnavailableError)
		return
	}

	if quotes[1].Error != nil {
		assert.Error(t, quotes[1].Error)
		assert.ErrorIs(t, quotes[1].Error, exchange.APIUnavailableError)
		return
	}

	assert.Equal(t, quotes[0].Currencies, "JPYUSD")
	assert.LessOrEqual(t, quotes[0].Rate, 89.5)
	assert.GreaterOrEqual(t, quotes[0].Rate, 89.0)
	assert.Equal(t, quotes[1].Currencies, "USDGBP")
	assert.LessOrEqual(t, quotes[1].Rate, 68.5)
	assert.GreaterOrEqual(t, quotes[1].Rate, 68.0)
}

func TestInvalidCurrency(t *testing.T) {
	c := New("fake-api-key")

	rate, err := c.GetRate("US", "GBP")
	assert.Equal(t, rate, 0.0)
	assert.Error(t, err)
	var invalidError *exchange.InvalidCurrencyError
	assert.ErrorAs(t, err, &invalidError)
}
