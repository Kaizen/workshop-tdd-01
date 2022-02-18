package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Note: I am going to ignore the fact you can do t.Run()

// So in Go you must start with "Test"
// so you can't do naming like "<verb>Should<action>" or "given<something>when<action>then<expectation>"

func TestIsSame(t *testing.T) {
	assert.Equal(t, 4, 3)
}
