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
