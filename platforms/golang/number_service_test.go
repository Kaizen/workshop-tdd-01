package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {

	t.Run("can be mocked", func(t *testing.T) {
		mock := NumberServiceMock{}
		mock.GenerateFunc = func() int { return 4 }

		actual := mock.Generate()

		assert.Equal(t, 4, actual)
		assert.Equal(t, 1, len(mock.GenerateCalls()))
	})
}
