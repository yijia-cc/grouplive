package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	subtraction := Subtract(6, 3)
	assert.Equal(t, 3, subtraction)
}
