package daraja_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SirWaithaka/payments/daraja"
)

func TestCache_Get(t *testing.T) {

	t.Run("test that empty string is returned when cache is empty", func(t *testing.T) {
		cache := daraja.NewCache[string]()

		assert.Empty(t, cache.Get())
	})

	t.Run("test that empty string is returned when cache is expired", func(t *testing.T) {
		cache := daraja.NewCache[string]()
		// set cache expiry to 10 seconds ago
		cache.Set("fake_value", time.Now().Add(-time.Second*10))

		assert.Empty(t, cache.Get())
	})

	t.Run("test that correct value is returned when cache is not expired or empty", func(t *testing.T) {
		cache := daraja.NewCache[string]()
		// set cache expiry to 10 seconds from now
		value := "fake_value"
		cache.Set(value, time.Now().Add(time.Second*10))

		assert.Equal(t, value, cache.Get())
	})
}
