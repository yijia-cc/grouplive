package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	subtraction := Subtract(2, 3)
	assert.Equal(t, -1, subtraction)
}
