package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRateSuccessWithCorrectParams(t *testing.T) {
	rate, err := GetRate("JPY", "USD")

	// If there is an error, that's okay, just assert that it's
	// the one we expect to see from the random failure generator
	if err != nil {
		assert.Equal(t, 0.0, rate)
		assert.Error(t, err)
		assert.ErrorIs(t, err, APIUnavailableError)
		return
	}

	assert.LessOrEqual(t, rate, 89.5)
	assert.GreaterOrEqual(t, rate, 89.0)
}
