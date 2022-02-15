package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {

	t.Run("returns 4", func(t *testing.T) {
		foo := foo()
		assert.Equal(t, 4, foo)
	})
}

func TestFetchAllCards(t *testing.T) {

	t.Run("returns data", func(t *testing.T) {
		data, err := fetchAllCards()

		assert.NoError(t, err)
		assert.Equal(t, 17449, len(data))
		assert.Equal(t, map[string]interface{}{
			"BasicType":        float64(4),
			"Cmc":              float64(5),
			"ColorIdentity":    float64(1),
			"ImageRelativeUrl": "/cards/normal/en/mir/51.jpg?1496448326",
			"Legality":         float64(408),
			"ManaCost":         "{4}{W}",
			"MultiverseId":     float64(0),
			"Name":             "Zuberi, Golden Feather",
			"Price":            float64(0.1),
			"Rating":           float64(0),
			"Types":            "Legendary Creature — Griffin",
		}, data[0])
	})
}
